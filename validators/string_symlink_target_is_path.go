package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringSymlinkTargetIsPathError is a function that defines error message returned by StringSymlinkTargetIsPath validator.
// nolint: gochecknoglobals
var StringSymlinkTargetIsPathError = func(v *StringSymlinkTargetIsPath) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("symlink's '%s' target is not an existing path", v.Field)
}

// StringSymlinkTargetIsPath is a validator object.
// Validate adds an error if the Field is a symlink and it's target is not an existing path.
type StringSymlinkTargetIsPath struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a symlink and it's target is not an existing path.
func (v *StringSymlinkTargetIsPath) Validate(e *validator.Errors) {

	target, err := os.Readlink(v.Field)
	if err != nil {
		return
	}

	if _, err := os.Stat(target); !os.IsNotExist(err) {
		return
	}

	e.Add(v.Name, StringSymlinkTargetIsPathError(v))
}

// SetField sets validator field.
func (v *StringSymlinkTargetIsPath) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringSymlinkTargetIsPath) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}