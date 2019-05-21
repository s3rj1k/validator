package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathWithoutSetUIDError is a function that defines error message returned by StringIsPathWithoutSetUID validator.
// nolint: gochecknoglobals
var StringIsPathWithoutSetUIDError = func(v *StringIsPathWithoutSetUID) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not an existing path or is an existing path with SetGID flag", v.Field)
}

// StringIsPathWithoutSetUID is a validator object.
type StringIsPathWithoutSetUID struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not an existing path or is an existing path with SetGID flag.
// If Field is a symlink, the symlink's target will be assessed.
func (v *StringIsPathWithoutSetUID) Validate(e *validator.Errors) {
	if Exists(v.Field) && !isFileWithMode(v.Field, os.ModeSetuid) {
		return
	}

	e.Add(v.Name, StringIsPathWithoutSetUIDError(v))
}

// SetField sets validator field.
func (v *StringIsPathWithoutSetUID) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPathWithoutSetUID) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
