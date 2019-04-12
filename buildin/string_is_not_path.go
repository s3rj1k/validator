package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsNotPathError is a function that defines error message returned by StringIsNotPath validator.
// nolint: gochecknoglobals
var StringIsNotPathError = func(v *StringIsNotPath) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("path '%s' must not exist", v.Field)
}

// StringIsNotPath is a validator object.
// Validate adds an error if the Field is an existing path.
type StringIsNotPath struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is an existing path.
func (v *StringIsNotPath) Validate(e *validator.Errors) {
	if _, err := os.Stat(v.Field); os.IsNotExist(err) {
		return
	}

	e.Add(v.Name, StringIsNotPathError(v))
}

// SetField sets validator field.
func (v *StringIsNotPath) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNotPath) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
