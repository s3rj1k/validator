package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsBase64Error is a function that defines error message returned by StringIsBase64 validator.
// nolint: gochecknoglobals
var StringIsBase64Error = func(v *StringIsBase64) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be base64 encoded", v.Field)
}

// StringIsBase64 is a validator object.
// Validate adds an error if the Field is not base64 encoded. Empty string is valid.
type StringIsBase64 struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not base64 encoded. Empty string is valid.
func (v *StringIsBase64) Validate(e *validator.Errors) {

	if isNullString(v.Field) {
		return
	}

	// base64 is valid
	if rxBase64.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringIsBase64Error(v))
}

// SetField sets validator field.
func (v *StringIsBase64) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsBase64) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
