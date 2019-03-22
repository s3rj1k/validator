package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringHasLowerCaseError is a function that defines error message returned by StringHasLowerCase validator.
// nolint: gochecknoglobals
var StringHasLowerCaseError = func(v *StringHasLowerCase) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must contain at least 1 lowercase", v.Field)
}

// StringHasLowerCase is a validator object
type StringHasLowerCase struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field has not lowercased letters. Empty string is valid.
func (v *StringHasLowerCase) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	if rxHasLowerCase.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringHasLowerCaseError(v))
}

// SetField sets validator field.
func (v *StringHasLowerCase) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasLowerCase) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
