package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsAlphaError is a function that defines error message returned by StringIsAlpha validator.
// nolint: gochecknoglobals
var StringIsAlphaError = func(v *StringIsAlpha) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must contain only letters", v.Field)
}

// StringIsAlpha is a validator object.
// Validate adds an error if the Field contains anything except for latin letters.
// Empty string is valid.
type StringIsAlpha struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field contains anything except for latin letters.
// Empty string is valid.
func (v *StringIsAlpha) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// alpha is valid, not trimming spaces
	if rxAlpha.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringIsAlphaError(v))
}

// SetField sets validator field.
func (v *StringIsAlpha) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsAlpha) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
