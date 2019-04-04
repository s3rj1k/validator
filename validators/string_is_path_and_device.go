package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathAndDeviceError is a function that defines error message returned by StringIsPathAndDevice validator.
// nolint: gochecknoglobals
var StringIsPathAndDeviceError = func(v *StringIsPathAndDevice) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not an existing path or is an existing path without Device mode", v.Field)
}

// StringIsPathAndDevice is a validator object.
// Validate adds an error if the Field is not an existing path or is an existing path without Device mode.
type StringIsPathAndDevice struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not an existing path or is an existing path without Device mode.
func (v *StringIsPathAndDevice) Validate(e *validator.Errors) {

	if Exists(v.Field) && isFileWithMode(v.Field, os.ModeDevice) {
		return
	}

	e.Add(v.Name, StringIsPathAndDeviceError(v))
}

// SetField sets validator field.
func (v *StringIsPathAndDevice) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPathAndDevice) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
