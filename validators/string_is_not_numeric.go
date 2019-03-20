package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsNotNumericError is a function that defines error message returned by StringIsNotNumeric validator.
// nolint: gochecknoglobals
var StringIsNotNumericError = func(v *StringIsNotNumeric) string {
	return fmt.Sprintf("'%s' is numeric", v.Field)
}

// StringIsNotNumeric is a validator object.
type StringIsNotNumeric struct {
	Name  string
	Field string
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
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}