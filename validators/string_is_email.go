package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsEmailError is a function that defines error message returned by StringIsEmail validator.
// nolint: gochecknoglobals
var StringIsEmailError = func(v *StringIsEmail) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' does not match an email format", v.Field)
}

// StringIsEmail is a validator object.
// Validate adds an error if the Field does not match email regexp. See Email const.
type StringIsEmail struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field does not match email regexp. See Email const.
func (v *StringIsEmail) Validate(e *validator.Errors) {
	if rxEmail.Match([]byte(v.Field)) {
		return
	}

	e.Add(v.Name, StringIsEmailError(v))

}

// SetField sets validator field.
func (v *StringIsEmail) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsEmail) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
