package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathAndNamedPipeError is a function that defines error message returned by StringIsPathAndNamedPipe validator.
// nolint: gochecknoglobals
var StringIsPathAndNamedPipeError = func(v *StringIsPathAndNamedPipe) string {
	return fmt.Sprintf("'%s' is not an existing path or is an existing path without NamedPipe mode", v.Field)
}

// StringIsPathAndNamedPipe is a validator object.
// Validate adds an error if the Field is not an existing path or is an existing path without NamedPipe mode.
type StringIsPathAndNamedPipe struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not an existing path or is an existing path without NamedPipe mode.
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
