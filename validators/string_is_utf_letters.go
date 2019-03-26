package validators

import (
	"fmt"
	"unicode"

	"github.com/s3rj1k/validator"
)

// StringIsUTFLettersError is a function that defines error message returned by StringIsUTFLetters validator.
// nolint: gochecknoglobals
var StringIsUTFLettersError = func(v *StringIsUTFLetters) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must contain only unicode letter characters", v.Field)
}

// StringIsUTFLetters is a validator object.
// Validate adds an error if the Field contains anything except unicode letters (category L)
// Similar to StringIsAlpha but for all languages. Empty string is valid.
type StringIsUTFLetters struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field contains anything except unicode letters (category L)
// Similar to StringIsAlpha but for all languages. Empty string is valid.
func (v *StringIsUTFLetters) Validate(e *validator.Errors) {
	var badRune bool

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	// checking each rune
	for _, c := range v.Field {
		if !unicode.IsLetter(c) {
			badRune = true
			break
		}
	}

	if badRune {
		e.Add(v.Name, StringIsUTFLettersError(v))
	}
}

// SetField sets validator field.
func (v *StringIsUTFLetters) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsUTFLetters) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
