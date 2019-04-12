package buildin

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsUpperCaseError is a function that defines error message returned by StringIsUpperCase validator.
// nolint: gochecknoglobals
var StringIsUpperCaseError = func(v *StringIsUpperCase) string {

	if len(v.Message) > 0 {
		return v.Message
	}
	return fmt.Sprintf("'%s' must be uppercased", v.Field)
}

// StringIsUpperCase is a validator object.
// Validate adds an error if the Field is not uppercased. Empty string is valid.
type StringIsUpperCase struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not uppercased. Empty string is valid.
func (v *StringIsUpperCase) Validate(e *validator.Errors) {

	// null string is valid
	if isNullString(v.Field) {
		return
	}

	if v.Field == strings.ToUpper(v.Field) {
		return
	}

	e.Add(v.Name, StringIsUpperCaseError(v))
}

// SetField sets validator field.
func (v *StringIsUpperCase) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsUpperCase) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
