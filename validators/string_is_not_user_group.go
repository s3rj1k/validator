package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsNotUserGroupError is a function that defines error message returned by StringIsNotUserGroup validator.
// nolint: gochecknoglobals
var StringIsNotUserGroupError = func(v *StringIsNotUserGroup) string {
	return fmt.Sprintf("'%s' is a user group", v.Name)
}

// StringIsNotUserGroup is a validator object.
type StringIsNotUserGroup struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is a user group.
func (v *StringIsNotUserGroup) Validate(e *validator.Errors) {

	if !IsGroupIsUserGroupOrWhitelisted(v.Field) {
		return
	}

	e.Add(v.Name, StringIsNotUserGroupError(v))
}

// SetField sets validator field.
func (v *StringIsNotUserGroup) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNotUserGroup) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
