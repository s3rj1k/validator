package validators

import (
	"fmt"
	"os/user"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsNotExistingGroupError is a function that defines error message returned by StringIsNotExistingGroup validator.
// nolint: gochecknoglobals
var StringIsNotExistingGroupError = func(v *StringIsNotExistingGroup) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' user group exists", v.Field)
}

// StringIsNotExistingGroup is a validator object.
type StringIsNotExistingGroup struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a user group that exists.
func (v *StringIsNotExistingGroup) Validate(e *validator.Errors) {

	_, err := user.LookupGroup(v.Field)
	if err != nil {
		return
	}

	e.Add(v.Name, StringIsNotExistingGroupError(v))
}

// SetField sets validator field.
func (v *StringIsNotExistingGroup) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNotExistingGroup) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
