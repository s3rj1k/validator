package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathAndRegularFileError is a function that defines error message returned by StringIsPathAndRegularFile validator.
// nolint: gochecknoglobals
var StringIsPathAndRegularFileError = func(v *StringIsPathAndRegularFile) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not an existing path or is an existing non-regular file", v.Field)
}

// StringIsPathAndRegularFile is a validator object.
// Validate adds an error if the Field is not an existing path or is an existing non-regular file.
// If Field is a symlink, the symlink's target will be assessed.
type StringIsPathAndRegularFile struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not an existing path or is an existing non-regular file.
// If Field is a symlink, the symlink's target will be assessed.
func (v *StringIsPathAndRegularFile) Validate(e *validator.Errors) {

	if Exists(v.Field) && isFileWithMode(v.Field, os.ModeType) {
		return
	}

	e.Add(v.Name, StringIsPathAndRegularFileError(v))
}

// SetField sets validator field.
func (v *StringIsPathAndRegularFile) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPathAndRegularFile) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
