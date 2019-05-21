package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathAndNotSocketError is a function that defines error message returned by StringIsPathAndNotSocket validator.
// nolint: gochecknoglobals
var StringIsPathAndNotSocketError = func(v *StringIsPathAndNotSocket) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not an existing path or is an existing path with Socket mode", v.Field)
}

// StringIsPathAndNotSocket is a validator object.
// Validate adds an error if the Field is not an existing path or is an existing path with Socket mode.
// If Field is a symlink, the symlink's target will be assessed.
type StringIsPathAndNotSocket struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not an existing path or is an existing path with Socket mode.
// If Field is a symlink, the symlink's target will be assessed.
func (v *StringIsPathAndNotSocket) Validate(e *validator.Errors) {
	if Exists(v.Field) && !isFileWithMode(v.Field, os.ModeSocket) {
		return
	}

	e.Add(v.Name, StringIsPathAndNotSocketError(v))
}

// SetField sets validator field.
func (v *StringIsPathAndNotSocket) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPathAndNotSocket) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
