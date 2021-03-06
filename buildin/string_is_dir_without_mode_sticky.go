package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsDirWithoutModeStickyError is a function that defines error message returned by StringIsDirWithoutModeSticky validator.
// nolint: gochecknoglobals
var StringIsDirWithoutModeStickyError = func(v *StringIsDirWithoutModeSticky) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not a dir or a dir with mode sticky", v.Field)
}

// StringIsDirWithoutModeSticky is a validator object.
// Validate adds an error if the Field is not a dir or a dir with mode sticky.
// If Field is a symlink, the symlink's target will be assessed.
type StringIsDirWithoutModeSticky struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a dir or a dir with mode sticky.
// If Field is a symlink, the symlink's target will be assessed.
func (v *StringIsDirWithoutModeSticky) Validate(e *validator.Errors) {
	if isDir(v.Field) && !isFileWithMode(v.Field, os.ModeSticky) {
		return
	}

	e.Add(v.Name, StringIsDirWithoutModeStickyError(v))
}

// SetField sets validator field.
func (v *StringIsDirWithoutModeSticky) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsDirWithoutModeSticky) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
