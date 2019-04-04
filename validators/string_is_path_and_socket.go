package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathAndSocketError is a function that defines error message returned by StringIsPathAndSocket validator.
// nolint: gochecknoglobals
var StringIsPathAndSocketError = func(v *StringIsPathAndSocket) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not an existing path or is an existing path without Socket mode", v.Field)
}

// StringIsPathAndSocket is a validator object.
// Validate adds an error if the Field is not an existing path or is an existing path without Socket mode.
// If Field is a symlink, the symlink's target will be assessed.
type StringIsPathAndSocket struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not an existing path or is an existing path without Socket mode.
// If Field is a symlink, the symlink's target will be assessed.
func (v *StringIsPathAndSocket) Validate(e *validator.Errors) {

	if Exists(v.Field) && isFileWithMode(v.Field, os.ModeSocket) {
		return
	}

	e.Add(v.Name, StringIsPathAndSocketError(v))
}

// SetField sets validator field.
func (v *StringIsPathAndSocket) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPathAndSocket) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
