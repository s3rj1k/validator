package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsValidUUIDv4Error is a function that defines error message returned by StringIsValidUUIDv4 validator.
// nolint: gochecknoglobals
var StringIsValidUUIDv4Error = func(v *StringIsValidUUIDv4) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not a valid UUIDv4", v.Field)
}

// StringIsValidUUIDv4 is a validator object.
// Validate adds an error if the Field is not a valid UUIDv4.
type StringIsValidUUIDv4 struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a valid UUIDv4.
func (v *StringIsValidUUIDv4) Validate(e *validator.Errors) {
	if rxUUIDv4.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringIsValidUUIDv4Error(v))
}

// SetField sets validator field.
func (v *StringIsValidUUIDv4) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsValidUUIDv4) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
