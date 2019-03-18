package validators

import (
	"fmt"
	"os/user"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsExistingGroupError is a function that defines error message returned by StringIsExistingGroup validator.
// nolint: gochecknoglobals
var StringIsExistingGroupError = func(v *StringIsExistingGroup) string {
	return fmt.Sprintf("'%s' user group does not exist", v.Field)
}

// StringIsExistingGroup is a validator object.
type StringIsExistingGroup struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is a user group that does not exist.
func (v *StringIsExistingGroup) Validate(e *validator.Errors) {

	_, err := user.LookupGroup(v.Field)
	if err == nil {
		return
	}

	e.Add(v.Name, StringIsExistingGroupError(v))
}

// SetField sets validator field.
func (v *StringIsExistingGroup) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsExistingGroup) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
