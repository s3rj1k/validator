package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathAndNotNamedPipeError is a function that defines error message returned by StringIsPathAndNotNamedPipe validator.
// nolint: gochecknoglobals
var StringIsPathAndNotNamedPipeError = func(v *StringIsPathAndNotNamedPipe) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not an existing path or is an existing path with NamedPipe mode", v.Field)
}

// StringIsPathAndNotNamedPipe is a validator object.
// Validate adds an error if the Field is not an existing path or is an existing path with NamedPipe mode.
// If Field is a symlink, the symlink's target will be assessed.
type StringIsPathAndNotNamedPipe struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not an existing path or is an existing path with NamedPipe mode.
// If Field is a symlink, the symlink's target will be assessed.
func (v *StringIsPathAndNotNamedPipe) Validate(e *validator.Errors) {
	if Exists(v.Field) && !isFileWithMode(v.Field, os.ModeNamedPipe) {
		return
	}

	e.Add(v.Name, StringIsPathAndNotNamedPipeError(v))
}

// SetField sets validator field.
func (v *StringIsPathAndNotNamedPipe) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPathAndNotNamedPipe) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
