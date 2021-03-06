package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsNullError is a function that defines error message returned by StringIsNull validator.
// nolint: gochecknoglobals
var StringIsNullError = func(v *StringIsNull) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be empty", v.Field)
}

// StringIsNull is a validator object.
// Validate adds an error if the Field is an empty string.
// Emptry string is defined as such with length 0.
// If you want to allow whitespaces - see StringIsPresent validator.
type StringIsNull struct {
	Name    string
	Field   string
	Message string
}

// isNullString is wrapper func
func isNullString(str string) bool {
	if len(str) == 0 { // nolint: megacheck
		return true
	}

	return false
}

// Validate adds an error if the Field is an empty string.
// Emptry string is defined as such with length 0.
// If you want to allow whitespaces - see StringIsPresent validator.
func (v *StringIsNull) Validate(e *validator.Errors) {
	if isNullString(v.Field) {
		return
	}

	e.Add(v.Name, StringIsNullError(v))
}

// SetField sets validator field.
func (v *StringIsNull) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNull) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
