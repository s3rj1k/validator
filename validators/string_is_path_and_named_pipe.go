package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathAndNamedPipeError is a function that defines error message returned by StringIsPathAndNamedPipe validator.
// nolint: gochecknoglobals
var StringIsPathAndNamedPipeError = func(v *StringIsPathAndNamedPipe) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not an existing path or is an existing path without NamedPipe mode", v.Field)
}

// StringIsPathAndNamedPipe is a validator object.
// Validate adds an error if the Field is not an existing path or is an existing path without NamedPipe mode.
// If Field is a symlink, the symlink's target will be assessed.
type StringIsPathAndNamedPipe struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not an existing path or is an existing path without NamedPipe mode.
// If Field is a symlink, the symlink's target will be assessed.
func (v *StringIsPathAndNamedPipe) Validate(e *validator.Errors) {

	if Exists(v.Field) && isFileWithMode(v.Field, os.ModeNamedPipe) {
		return
	}

	e.Add(v.Name, StringIsPathAndNamedPipeError(v))
}

// SetField sets validator field.
func (v *StringIsPathAndNamedPipe) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPathAndNamedPipe) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
