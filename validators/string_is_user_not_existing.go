package validators

import (
	"fmt"
	"os/user"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsUserNotExistingError is a function that defines error message returned by StringIsUserNotExisting validator.
// nolint: gochecknoglobals
var StringIsUserNotExistingError = func(v *StringIsUserNotExisting) string {
	return fmt.Sprintf("'%s' user exists", v.Name)
}

// StringIsUserNotExisting is a validator object.
type StringIsUserNotExisting struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is an existing user.
func (v *StringIsUserNotExisting) Validate(e *validator.Errors) {

	_, err := user.Lookup(v.Field)
	if err != nil {
		return
	}

	e.Add(v.Name, StringIsUserNotExistingError(v))
}

// SetField sets validator field.
func (v *StringIsUserNotExisting) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsUserNotExisting) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
