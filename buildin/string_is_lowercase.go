package buildin

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsLowerCaseError is a function that defines error message returned by StringIsLowerCase validator.
// nolint: gochecknoglobals
var StringIsLowerCaseError = func(v *StringIsLowerCase) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be lowercased", v.Field)
}

// StringIsLowerCase is a validator object.
// Validate adds an error if the Field is not lowercased. Empty string is valid.
type StringIsLowerCase struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not lowercased. Empty string is valid.
func (v *StringIsLowerCase) Validate(e *validator.Errors) {
	// null string is valid
	if isNullString(v.Field) {
		return
	}

	if v.Field == strings.ToLower(v.Field) {
		return
	}

	e.Add(v.Name, StringIsLowerCaseError(v))
}

// SetField sets validator field.
func (v *StringIsLowerCase) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsLowerCase) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
