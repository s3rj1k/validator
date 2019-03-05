package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsHexadecimalError is a function that defines error message returned by StringIsHexadecimal validator.
// nolint: gochecknoglobals
var StringIsHexadecimalError = func(v *StringIsHexadecimal) string {
	return fmt.Sprintf("%s must be a hexadecimal number", v.Name)
}

// StringIsHexadecimal is a validator object.
type StringIsHexadecimal struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not in a hexadecimal format.
func (v *StringIsHexadecimal) Validate(e *validator.Errors) {

	if rxHexadecimal.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringIsHexadecimalError(v))
}

// SetField sets validator field.
func (v *StringIsHexadecimal) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsHexadecimal) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
