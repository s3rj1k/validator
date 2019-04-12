package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsNotNumericError is a function that defines error message returned by StringIsNotNumeric validator.
// nolint: gochecknoglobals
var StringIsNotNumericError = func(v *StringIsNotNumeric) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is numeric", v.Field)
}

// StringIsNotNumeric is a validator object.
// Validate adds an error if the Field is numeric.
type StringIsNotNumeric struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is numeric.
func (v *StringIsNotNumeric) Validate(e *validator.Errors) {

	if !rxNumeric.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringIsNotNumericError(v))
}

// SetField sets validator field.
func (v *StringIsNotNumeric) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNotNumeric) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
