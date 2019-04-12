package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsPrintableASCIIError is a function that defines error message returned by StringIsPrintableASCII validator.
// nolint: gochecknoglobals
var StringIsPrintableASCIIError = func(v *StringIsPrintableASCII) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must contain printable ASCII chars only", v.Field)
}

// StringIsPrintableASCII is a validator object.
// Validate adds an error if the Field contains anything except for printable ASCII characters.
// Empty string is valid.
type StringIsPrintableASCII struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field contains anything except for printable ASCII characters.
// Empty string is valid.
func (v *StringIsPrintableASCII) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// ASCII is valid
	if rxPrintableASCII.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringIsPrintableASCIIError(v))
}

// SetField sets validator field.
func (v *StringIsPrintableASCII) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPrintableASCII) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
