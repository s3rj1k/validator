package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsNotRegularUserError is a function that defines error message returned by StringIsNotRegularUser validator.
// nolint: gochecknoglobals
var StringIsNotRegularUserError = func(v *StringIsNotRegularUser) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is a regular user", v.Field)
}

// StringIsNotRegularUser is a validator object.
type StringIsNotRegularUser struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a regular user.
func (v *StringIsNotRegularUser) Validate(e *validator.Errors) {
	if !IsUserIsRegularUserOrWhitelisted(v.Field) {
		return
	}

	e.Add(v.Name, StringIsNotRegularUserError(v))
}

// SetField sets validator field.
func (v *StringIsNotRegularUser) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNotRegularUser) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
