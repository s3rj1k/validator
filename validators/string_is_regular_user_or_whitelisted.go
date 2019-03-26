package validators

import (
	"fmt"
	"os/user"
	"strconv"

	"github.com/s3rj1k/validator"
)

// StringIsRegularUserOrWhitelistedError is a function that defines error message returned by StringIsRegularUserOrWhitelisted validator.
// nolint: gochecknoglobals
var StringIsRegularUserOrWhitelistedError = func(v *StringIsRegularUserOrWhitelisted) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not a regular user or whitelisted", v.Field)
}

// StringIsRegularUserOrWhitelisted is a validator object.
// Validate adds an error if the Field is not a regular user or whitelisted.
type StringIsRegularUserOrWhitelisted struct {
	Name      string
	Field     string
	Whitelist []string
	Message   string
}

// Validate adds an error if the Field is not a regular user or whitelisted.
func (v *StringIsRegularUserOrWhitelisted) Validate(e *validator.Errors) {

	if IsUserIsRegularUserOrWhitelisted(v.Field, v.Whitelist...) {
		return
	}

	e.Add(v.Name, StringIsRegularUserOrWhitelistedError(v))
}

// SetField sets validator field.
func (v *StringIsRegularUserOrWhitelisted) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsRegularUserOrWhitelisted) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}

// IsUserIsRegularUserOrWhitelisted checkes if user is regular (non-system or unclassified) user.
func IsUserIsRegularUserOrWhitelisted(name string, whitelist ...string) bool {

	// check whitelist of groups
	for _, el := range whitelist {
		if el == name {
			return true
		}
	}

	user, err := user.Lookup(name)
	if err != nil {
		return false // fail on lookup error
	}

	uid, err := strconv.ParseUint(user.Uid, 10, 32)
	if err != nil {
		return false // fail on parse error
	}

	minGID, maxGiD := ReadUserUIDRange(LoginDefsPath)

	if uid < minGID {
		return false // user not in lower range
	}

	if uid > maxGiD {
		return false // user not in upper range
	}

	return true
}
