package buildin

import (
	"fmt"
	"os/user"
	"strconv"

	"github.com/s3rj1k/validator"
)

// StringIsUserGroupOrWhitelistedError is a function that defines error message returned by StringIsUserGroupOrWhitelisted validator.
// nolint: gochecknoglobals
var StringIsUserGroupOrWhitelistedError = func(v *StringIsUserGroupOrWhitelisted) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not a user group or whitelisted", v.Field)
}

// StringIsUserGroupOrWhitelisted is a validator object.
// Validate adds an error if the Field is not a user group or whitelisted.
type StringIsUserGroupOrWhitelisted struct {
	Name      string
	Field     string
	Whitelist []string
	Message   string
}

// Validate adds an error if the Field is not a user group or whitelisted.
func (v *StringIsUserGroupOrWhitelisted) Validate(e *validator.Errors) {

	if IsGroupIsUserGroupOrWhitelisted(v.Field, v.Whitelist...) {
		return
	}

	e.Add(v.Name, StringIsUserGroupOrWhitelistedError(v))
}

// SetField sets validator field.
func (v *StringIsUserGroupOrWhitelisted) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsUserGroupOrWhitelisted) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}

// IsGroupIsUserGroupOrWhitelisted checkes if name is an allowed name of group
func IsGroupIsUserGroupOrWhitelisted(name string, whitelist ...string) bool {

	// check whitelist of groups
	for _, el := range whitelist {
		if el == name {
			return true
		}
	}

	group, err := user.LookupGroup(name)
	if err != nil {
		return false // fail on lookup error
	}

	gid, err := strconv.ParseUint(group.Gid, 10, 32)
	if err != nil {
		return false // fail on parse error
	}

	minGID, maxGiD := ReadUserGIDRange(LoginDefsPath)

	if gid < minGID {
		return false // group not in lower range
	}

	if gid > maxGiD {
		return false // group not in upper range
	}

	return true
}
