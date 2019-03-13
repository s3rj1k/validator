package validators

import (
	"fmt"
	"os/user"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsUserExistingError is a function that defines error message returned by StringIsUserExisting validator.
// nolint: gochecknoglobals
var StringIsUserExistingError = func(v *StringIsUserExisting) string {
	return fmt.Sprintf("'%s' user does not exist", v.Name)
}

// StringIsUserExisting is a validator object.
type StringIsUserExisting struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is a user that does not exist.
func (v *StringIsUserExisting) Validate(e *validator.Errors) {

	_, err := user.Lookup(v.Field)
	if err == nil {
		return
	}

	e.Add(v.Name, StringIsUserExistingError(v))
}

// SetField sets validator field.
func (v *StringIsUserExisting) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsUserExisting) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
