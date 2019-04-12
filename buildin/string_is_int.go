package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsIntError is a function that defines error message returned by StringIsInt validator.
// nolint: gochecknoglobals
var StringIsIntError = func(v *StringIsInt) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be an integer", v.Field)
}

// StringIsInt is a validator object.
// Validate adds an error if the Field is not an integer.
// Leading sign is allowed. Empty string is valid.
type StringIsInt struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not an integer.
// Leading sign is allowed. Empty string is valid.
func (v *StringIsInt) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// Int is valid
	if rxInteger.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringIsIntError(v))
}

// SetField sets validator field.
func (v *StringIsInt) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsInt) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
