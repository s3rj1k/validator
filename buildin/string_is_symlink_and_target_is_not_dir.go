package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsSymlinkAndTargetIsNotDirError is a function that defines error message returned by StringIsSymlinkAndTargetIsNotDir validator.
// nolint: gochecknoglobals
var StringIsSymlinkAndTargetIsNotDirError = func(v *StringIsSymlinkAndTargetIsNotDir) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("symlink's '%s' target is a directory", v.Field)
}

// StringIsSymlinkAndTargetIsNotDir is a validator object.
// Validate adds an error if the Field is a symlink and it's target is a directory.
type StringIsSymlinkAndTargetIsNotDir struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a symlink and it's target is a directory.
func (v *StringIsSymlinkAndTargetIsNotDir) Validate(e *validator.Errors) {

	target, err := os.Readlink(v.Field)
	if err != nil {
		return
	}

	fi, err := os.Stat(target)
	if err != nil {
		return
	}

	if mode := fi.Mode(); !mode.IsDir() {
		return
	}

	e.Add(v.Name, StringIsSymlinkAndTargetIsNotDirError(v))
}

// SetField sets validator field.
func (v *StringIsSymlinkAndTargetIsNotDir) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsSymlinkAndTargetIsNotDir) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
