package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsPrintableASCIIError is a function that defines error message returned by StringIsPrintableASCII validator.
// nolint: gochecknoglobals
var StringIsPrintableASCIIError = func(v *StringIsPrintableASCII) string {
	return fmt.Sprintf("'%s' must contain printable ASCII chars only", v.Field)
}

// StringIsPrintableASCII is a validator object.
type StringIsPrintableASCII struct {
	Name  string
	Field string
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
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
