package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsSymlinkAndTargetIsPathError is a function that defines error message returned by StringIsSymlinkAndTargetIsPath validator.
// nolint: gochecknoglobals
var StringIsSymlinkAndTargetIsPathError = func(v *StringIsSymlinkAndTargetIsPath) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("symlink's '%s' target is not an existing path", v.Field)
}

// StringIsSymlinkAndTargetIsPath is a validator object.
// Validate adds an error if the Field is a symlink and it's target is not an existing path.
type StringIsSymlinkAndTargetIsPath struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a symlink and it's target is not an existing path.
func (v *StringIsSymlinkAndTargetIsPath) Validate(e *validator.Errors) {

	target, err := os.Readlink(v.Field)
	if err != nil {
		return
	}

	if _, err := os.Stat(target); !os.IsNotExist(err) {
		return
	}

	e.Add(v.Name, StringIsSymlinkAndTargetIsPathError(v))
}

// SetField sets validator field.
func (v *StringIsSymlinkAndTargetIsPath) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsSymlinkAndTargetIsPath) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
