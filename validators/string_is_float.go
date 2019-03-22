package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsFloatError is a function that defines error message returned by StringIsFloat validator.
// nolint: gochecknoglobals
var StringIsFloatError = func(v *StringIsFloat) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be a float", v.Field)
}

// StringIsFloat is a validator object.
type StringIsFloat struct {
	Name    string
	Field   string
	Message string
}

// Validate add an error if the Field is not a float. Empty string is valid.
func (v *StringIsFloat) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// float is valid
	if rxFloat.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringIsFloatError(v))
}

// SetField sets validator field.
func (v *StringIsFloat) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsFloat) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
