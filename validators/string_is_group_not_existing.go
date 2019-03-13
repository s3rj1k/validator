package validators

import (
	"fmt"
	"os/user"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsGroupNotExistingError is a function that defines error message returned by StringIsGroupNotExisting validator.
// nolint: gochecknoglobals
var StringIsGroupNotExistingError = func(v *StringIsGroupNotExisting) string {
	return fmt.Sprintf("'%s' user group exists", v.Name)
}

// StringIsGroupNotExisting is a validator object.
type StringIsGroupNotExisting struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is a user group that exists.
func (v *StringIsGroupNotExisting) Validate(e *validator.Errors) {

	_, err := user.LookupGroup(v.Field)
	if err != nil {
		return
	}

	e.Add(v.Name, StringIsGroupNotExistingError(v))
}

// SetField sets validator field.
func (v *StringIsGroupNotExisting) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsGroupNotExisting) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
