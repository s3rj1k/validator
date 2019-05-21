package buildin

import (
	"fmt"
	"os/user"

	"github.com/s3rj1k/validator"
)

// StringIsNotExistingUserError is a function that defines error message returned by StringIsNotExistingUser validator.
// nolint: gochecknoglobals
var StringIsNotExistingUserError = func(v *StringIsNotExistingUser) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' user exists", v.Field)
}

// StringIsNotExistingUser is a validator object.
// Validate adds an error if the Field is an existing user.
type StringIsNotExistingUser struct {
	Name    string
	Field   string
	Message string
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
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
