package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsASCIIError is a function that defines error message returned by StringIsASCII validator.
// nolint: gochecknoglobals
var StringIsASCIIError = func(v *StringIsASCII) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must contain ASCII chars only", v.Field)
}

// StringIsASCII is a validator object.
// Validate adds an error if the Field contains anything except for ASCII characters.
// Empty string is valid.
type StringIsASCII struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field contains anything except for ASCII characters.
// Empty string is valid.
func (v *StringIsASCII) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// ASCII is valid
	if rxASCII.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringIsASCIIError(v))
}

// SetField sets validator field.
func (v *StringIsASCII) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsASCII) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
