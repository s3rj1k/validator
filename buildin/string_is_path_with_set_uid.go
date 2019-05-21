package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathWithSetUIDError is a function that defines error message returned by StringIsPathWithSetUID validator.
// nolint: gochecknoglobals
var StringIsPathWithSetUIDError = func(v *StringIsPathWithSetUID) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not an existing path or is an existing path without SetUID flag", v.Field)
}

// StringIsPathWithSetUID is a validator object.
// Validate adds an error if the Field is not an existing path or is an existing path without SetUID flag.
// If Field is a symlink, the symlink's target will be assessed.
type StringIsPathWithSetUID struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not an existing path or is an existing path without SetUID flag.
// If Field is a symlink, the symlink's target will be assessed.
func (v *StringIsPathWithSetUID) Validate(e *validator.Errors) {
	if Exists(v.Field) && isFileWithMode(v.Field, os.ModeSetuid) {
		return
	}

	e.Add(v.Name, StringIsPathWithSetUIDError(v))
}

// SetField sets validator field.
func (v *StringIsPathWithSetUID) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPathWithSetUID) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
