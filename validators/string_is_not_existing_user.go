package validators

import (
	"fmt"
	"os/user"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsNotExistingUserError is a function that defines error message returned by StringIsNotExistingUser validator.
// nolint: gochecknoglobals
var StringIsNotExistingUserError = func(v *StringIsNotExistingUser) string {
	return fmt.Sprintf("'%s' user exists", v.Name)
}

// StringIsNotExistingUser is a validator object.
type StringIsNotExistingUser struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is an existing user.
func (v *StringIsNotExistingUser) Validate(e *validator.Errors) {

	_, err := user.Lookup(v.Field)
	if err != nil {
		return
	}

	e.Add(v.Name, StringIsNotExistingUserError(v))
}

// SetField sets validator field.
func (v *StringIsNotExistingUser) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNotExistingUser) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}