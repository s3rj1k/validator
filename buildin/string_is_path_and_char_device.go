package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathAndCharDeviceError is a function that defines error message returned by StringIsPathAndCharDevice validator.
// nolint: gochecknoglobals
var StringIsPathAndCharDeviceError = func(v *StringIsPathAndCharDevice) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not an existing path or is an existing path without CharDevice mode", v.Field)
}

// StringIsPathAndCharDevice is a validator object.
// Validate adds an error if the Field is not an existing path or is an existing path without CharDevice mode.
// If Field is a symlink, the symlink's target will be assessed.
type StringIsPathAndCharDevice struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not an existing path or is an existing path without CharDevice mode.
// If Field is a symlink, the symlink's target will be assessed.
func (v *StringIsPathAndCharDevice) Validate(e *validator.Errors) {
	if Exists(v.Field) && isFileWithMode(v.Field, os.ModeCharDevice) {
		return
	}

	e.Add(v.Name, StringIsPathAndCharDeviceError(v))
}

// SetField sets validator field.
func (v *StringIsPathAndCharDevice) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPathAndCharDevice) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
