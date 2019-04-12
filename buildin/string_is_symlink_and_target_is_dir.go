package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsSymlinkAndTargetIsDirError is a function that defines error message returned by StringIsSymlinkAndTargetIsDir validator.
// nolint: gochecknoglobals
var StringIsSymlinkAndTargetIsDirError = func(v *StringIsSymlinkAndTargetIsDir) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("symlink's '%s' target is not a directory", v.Field)
}

// StringIsSymlinkAndTargetIsDir is a validator object.
// Validate adds an error if the Field is a symlink and it's target is not a directory.
type StringIsSymlinkAndTargetIsDir struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a symlink and it's target is not a directory.
func (v *StringIsSymlinkAndTargetIsDir) Validate(e *validator.Errors) {

	target, err := os.Readlink(v.Field)
	if err != nil {
		return
	}

	if fi, err := os.Stat(target); !os.IsNotExist(err) {
		if mode := fi.Mode(); mode.IsDir() {
			return
		}
	}

	e.Add(v.Name, StringIsSymlinkAndTargetIsDirError(v))
}

// SetField sets validator field.
func (v *StringIsSymlinkAndTargetIsDir) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsSymlinkAndTargetIsDir) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
