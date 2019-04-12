package buildin

import (
	"fmt"
	"os/user"

	"github.com/s3rj1k/validator"
)

// StringIsExistingGroupError is a function that defines error message returned by StringIsExistingGroup validator.
// nolint: gochecknoglobals
var StringIsExistingGroupError = func(v *StringIsExistingGroup) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' user group does not exist", v.Field)
}

// StringIsExistingGroup is a validator object.
// Validate adds an error if the Field is a user group that does not exist.
type StringIsExistingGroup struct {
	Name    string
	Field   string
	Message string
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
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
