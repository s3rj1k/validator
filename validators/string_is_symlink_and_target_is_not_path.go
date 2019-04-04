package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsSymlinkAndTargetIsNotPathError is a function that defines error message returned by StringIsSymlinkAndTargetIsNotPath validator.
// nolint: gochecknoglobals
var StringIsSymlinkAndTargetIsNotPathError = func(v *StringIsSymlinkAndTargetIsNotPath) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("symlink's '%s' target is an existing path", v.Field)
}

// StringIsSymlinkAndTargetIsNotPath is a validator object.
// Validate adds an error if the Field is a symlink and it's target is an existing path.
type StringIsSymlinkAndTargetIsNotPath struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a symlink and it's target is an existing path.
func (v *StringIsSymlinkAndTargetIsNotPath) Validate(e *validator.Errors) {

	target, err := os.Readlink(v.Field)
	if err != nil {
		return
	}

	if _, err := os.Stat(target); os.IsNotExist(err) {
		return
	}

	e.Add(v.Name, StringIsSymlinkAndTargetIsNotPathError(v))
}

// SetField sets validator field.
func (v *StringIsSymlinkAndTargetIsNotPath) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsSymlinkAndTargetIsNotPath) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
