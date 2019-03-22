package validators

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/s3rj1k/validator"
)

// StringIsUTFLetterNumError is a function that defines error message returned by StringIsUTFLetterNum validator.
// nolint: gochecknoglobals
var StringIsUTFLetterNumError = func(v *StringIsUTFLetterNum) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must contain only unicode letter/number characters", v.Field)
}

// StringIsUTFLetterNum is a validator object.
type StringIsUTFLetterNum struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field contains anything except unicode letters/numbers (category L/N).
// Empty string is valid.
func (v *StringIsUTFLetterNum) Validate(e *validator.Errors) {
	var badRune bool

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// checking each rune
	for _, c := range v.Field {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
			badRune = true
			break
		}
	}

	if badRune {
		e.Add(v.Name, StringIsUTFLetterNumError(v))
	}
}

// SetField sets validator field.
func (v *StringIsUTFLetterNum) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsUTFLetterNum) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
