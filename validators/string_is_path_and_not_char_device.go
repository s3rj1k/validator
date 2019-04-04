package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathAndNotCharDeviceError is a function that defines error message returned by StringIsPathAndNotCharDevice validator.
// nolint: gochecknoglobals
var StringIsPathAndNotCharDeviceError = func(v *StringIsPathAndNotCharDevice) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not an existing path or is an existing path with CharDevice mode", v.Field)
}

// StringIsPathAndNotCharDevice is a validator object.
// Validate adds an error if the Field is not an existing path or is an existing path with CharDevice mode.
type StringIsPathAndNotCharDevice struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not an existing path or is an existing path with CharDevice mode.
func (v *StringIsPathAndNotCharDevice) Validate(e *validator.Errors) {

	if Exists(v.Field) && !isFileWithMode(v.Field, os.ModeCharDevice) {
		return
	}

	e.Add(v.Name, StringIsPathAndNotCharDeviceError(v))
}

// SetField sets validator field.
func (v *StringIsPathAndNotCharDevice) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPathAndNotCharDevice) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
