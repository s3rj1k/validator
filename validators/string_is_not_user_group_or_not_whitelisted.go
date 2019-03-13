package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsNotUserGroupOrNotWhitelistedError is a function that defines error message returned by StringIsNotUserGroupOrNotWhitelisted validator.
// nolint: gochecknoglobals
var StringIsNotUserGroupOrNotWhitelistedError = func(v *StringIsNotUserGroupOrNotWhitelisted) string {
	return fmt.Sprintf("'%s' is a user group or whitelisted", v.Name)
}

// StringIsNotUserGroupOrNotWhitelisted is a validator object.
type StringIsNotUserGroupOrNotWhitelisted struct {
	Name      string
	Field     string
	Whitelist []string
}

// Validate adds an error if the Field is a user group or whitelisted
func (v *StringIsNotUserGroupOrNotWhitelisted) Validate(e *validator.Errors) {

	if !IsGroupIsUserGroupOrWhitelisted(v.Field, v.Whitelist...) {
		return
	}

	e.Add(v.Name, StringIsNotUserGroupOrNotWhitelistedError(v))
}

// SetField sets validator field.
func (v *StringIsNotUserGroupOrNotWhitelisted) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNotUserGroupOrNotWhitelisted) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
