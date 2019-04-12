package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathAndNotDeviceError is a function that defines error message returned by StringIsPathAndNotDevice validator.
// nolint: gochecknoglobals
var StringIsPathAndNotDeviceError = func(v *StringIsPathAndNotDevice) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not an existing path or is an existing path with Device mode", v.Field)
}

// StringIsPathAndNotDevice is a validator object.
// Validate adds an error if the Field is not an existing path or is an existing path with Device mode.
// If Field is a symlink, the symlink's target will be assessed.
type StringIsPathAndNotDevice struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not an existing path or is an existing path with Device mode.
// If Field is a symlink, the symlink's target will be assessed.
func (v *StringIsPathAndNotDevice) Validate(e *validator.Errors) {

	if Exists(v.Field) && !isFileWithMode(v.Field, os.ModeDevice) {
		return
	}

	e.Add(v.Name, StringIsPathAndNotDeviceError(v))
}

// SetField sets validator field.
func (v *StringIsPathAndNotDevice) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPathAndNotDevice) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
