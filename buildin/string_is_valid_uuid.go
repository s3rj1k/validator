package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsValidUUIDError is a function that defines error message returned by StringIsValidUUID validator.
// nolint: gochecknoglobals
var StringIsValidUUIDError = func(v *StringIsValidUUID) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not a valid UUID", v.Field)
}

// StringIsValidUUID is a validator object.
// Validate adds an error if the Field is not a valid UUID.
type StringIsValidUUID struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a valid UUID.
func (v *StringIsValidUUID) Validate(e *validator.Errors) {
	if rxUUID.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringIsValidUUIDError(v))
}

// SetField sets validator field.
func (v *StringIsValidUUID) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsValidUUID) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
