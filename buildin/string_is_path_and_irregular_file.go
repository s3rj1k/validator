package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathAndIrregularFileError is a function that defines error message returned by StringIsPathAndIrregularFile validator.
// nolint: gochecknoglobals
var StringIsPathAndIrregularFileError = func(v *StringIsPathAndIrregularFile) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not an existing path or is an existing path without Irregular mode", v.Field)
}

// StringIsPathAndIrregularFile is a validator object.
// Validate adds an error if the Field is not an existing path or is an existing path without Irregular mode.
// If Field is a symlink, the symlink's target will be assessed.
type StringIsPathAndIrregularFile struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not an existing path or is an existing path without Irregular mode.
// If Field is a symlink, the symlink's target will be assessed.
func (v *StringIsPathAndIrregularFile) Validate(e *validator.Errors) {

	if Exists(v.Field) && isFileWithMode(v.Field, os.ModeIrregular) {
		return
	}

	e.Add(v.Name, StringIsPathAndIrregularFileError(v))
}

// SetField sets validator field.
func (v *StringIsPathAndIrregularFile) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPathAndIrregularFile) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
