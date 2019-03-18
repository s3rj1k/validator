package validators

import (
	"fmt"
	"os/user"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsExistingUserError is a function that defines error message returned by StringIsExistingUser validator.
// nolint: gochecknoglobals
var StringIsExistingUserError = func(v *StringIsExistingUser) string {
	return fmt.Sprintf("'%s' user does not exist", v.Field)
}

// StringIsExistingUser is a validator object.
type StringIsExistingUser struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is a user that does not exist.
func (v *StringIsExistingUser) Validate(e *validator.Errors) {

	_, err := user.Lookup(v.Field)
	if err == nil {
		return
	}

	e.Add(v.Name, StringIsExistingUserError(v))
}

// SetField sets validator field.
func (v *StringIsExistingUser) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsExistingUser) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
