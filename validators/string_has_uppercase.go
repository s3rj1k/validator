package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringHasUpperCaseError is a function that defines error message returned by StringHasUpperCase validator.
// nolint: gochecknoglobals
var StringHasUpperCaseError = func(v *StringHasUpperCase) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must contain at least 1 uppercase", v.Field)
}

// StringHasUpperCase is a validator object.
// Validate adds an error if the Field has not uppercased letters. Empty string is valid.
type StringHasUpperCase struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field has not uppercased letters. Empty string is valid.
func (v *StringHasUpperCase) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	if rxHasUpperCase.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringHasUpperCaseError(v))
}

// SetField sets validator field.
func (v *StringHasUpperCase) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasUpperCase) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
