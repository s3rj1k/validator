package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsAlphaNumError is a function that defines error message returned by StringIsAlphaNum validator.
// nolint: gochecknoglobals
var StringIsAlphaNumError = func(v *StringIsAlphaNum) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must contain only numbers and/or letters", v.Field)
}

// StringIsAlphaNum is a validator object.
type StringIsAlphaNum struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field contains any symbols except for arabic numerals and latin letters.
// Empty string is valid.
func (v *StringIsAlphaNum) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// alphanum is valid, not trimming spaces
	if rxAlphanumeric.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringIsAlphaNumError(v))
}

// SetField sets validator field.
func (v *StringIsAlphaNum) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsAlphaNum) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
