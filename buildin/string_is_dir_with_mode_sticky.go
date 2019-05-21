package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsDirWithModeStickyError is a function that defines error message returned by StringIsDirWithModeSticky validator.
// nolint: gochecknoglobals
var StringIsDirWithModeStickyError = func(v *StringIsDirWithModeSticky) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not a dir or a dir without sticky mode", v.Field)
}

// StringIsDirWithModeSticky is a validator object.
// Validate adds an error if the Field is not a dir or a dir without sticky mode.
// If Field is a symlink, the symlink's target will be assessed.
type StringIsDirWithModeSticky struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a dir or a dir without sticky mode.
// If Field is a symlink, the symlink's target will be assessed.
func (v *StringIsDirWithModeSticky) Validate(e *validator.Errors) {
	if isDir(v.Field) && isFileWithMode(v.Field, os.ModeSticky) {
		return
	}

	e.Add(v.Name, StringIsDirWithModeStickyError(v))
}

// SetField sets validator field.
func (v *StringIsDirWithModeSticky) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsDirWithModeSticky) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
