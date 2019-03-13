package validators

import (
	"fmt"
	"os/user"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsGroupExistingError is a function that defines error message returned by StringIsGroupExisting validator.
// nolint: gochecknoglobals
var StringIsGroupExistingError = func(v *StringIsGroupExisting) string {
	return fmt.Sprintf("'%s' user group does not exist", v.Name)
}

// StringIsGroupExisting is a validator object.
type StringIsGroupExisting struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is a user group that does not exist.
func (v *StringIsGroupExisting) Validate(e *validator.Errors) {

	_, err := user.LookupGroup(v.Field)
	if err == nil {
		return
	}

	e.Add(v.Name, StringIsGroupExistingError(v))
}

// SetField sets validator field.
func (v *StringIsGroupExisting) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsGroupExisting) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
