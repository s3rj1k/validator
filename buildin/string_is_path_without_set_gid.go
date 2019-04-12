package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsPathWithoutSetGIDError is a function that defines error message returned by StringIsPathWithoutSetGID validator.
// nolint: gochecknoglobals
var StringIsPathWithoutSetGIDError = func(v *StringIsPathWithoutSetGID) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not an existing path or is an existing path with SetGID flag", v.Field)
}

// StringIsPathWithoutSetGID is a validator object.
// Validate adds an error if the Field is not an existing path or is an existing path with SetGID flag.
// If Field is a symlink, the symlink's target will be assessed.
type StringIsPathWithoutSetGID struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not an existing path or is an existing path with SetGID flag.
// If Field is a symlink, the symlink's target will be assessed.
func (v *StringIsPathWithoutSetGID) Validate(e *validator.Errors) {

	if Exists(v.Field) && !isFileWithMode(v.Field, os.ModeSetgid) {
		return
	}

	e.Add(v.Name, StringIsPathWithoutSetGIDError(v))
}

// SetField sets validator field.
func (v *StringIsPathWithoutSetGID) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPathWithoutSetGID) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
