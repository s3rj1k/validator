package validators

import (
	"fmt"
	"os/user"

	"github.com/s3rj1k/validator"
)

// StringIsExistingUserError is a function that defines error message returned by StringIsExistingUser validator.
// nolint: gochecknoglobals
var StringIsExistingUserError = func(v *StringIsExistingUser) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' user does not exist", v.Field)
}

// StringIsExistingUser is a validator object.
// Validate adds an error if the Field is a user that does not exist.
type StringIsExistingUser struct {
	Name    string
	Field   string
	Message string
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
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
