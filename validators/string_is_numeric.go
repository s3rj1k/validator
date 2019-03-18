package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsNumericError is a function that defines error message returned by StringIsNumeric validator.
// nolint: gochecknoglobals
var StringIsNumericError = func(v *StringIsNumeric) string {
	return fmt.Sprintf("'%s' must contain only numbers", v.Field)
}

// StringIsNumeric is a validator object.
type StringIsNumeric struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not numeric. Empty string is valid.
func (v *StringIsNumeric) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// alphanum is valid, not trimming spaces
	if rxNumeric.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringIsNumericError(v))
}

// SetField sets validator field.
func (v *StringIsNumeric) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNumeric) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
