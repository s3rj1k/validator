package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathWithSetGIDError is a function that defines error message returned by StringIsPathWithSetGID validator.
// nolint: gochecknoglobals
var StringIsPathWithSetGIDError = func(v *StringIsPathWithSetGID) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not an existing path or is an existing path without SetGID flag", v.Field)
}

// StringIsPathWithSetGID is a validator object.
// Validate adds an error if the Field is not an existing path or is an existing path without SetGID flag.
// If Field is a symlink, the symlink's target will be assessed.
type StringIsPathWithSetGID struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not an existing path or is an existing path without SetGID flag.
// If Field is a symlink, the symlink's target will be assessed.
func (v *StringIsPathWithSetGID) Validate(e *validator.Errors) {

	if Exists(v.Field) && isFileWithMode(v.Field, os.ModeSetgid) {
		return
	}

	e.Add(v.Name, StringIsPathWithSetGIDError(v))
}

// SetField sets validator field.
func (v *StringIsPathWithSetGID) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPathWithSetGID) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
