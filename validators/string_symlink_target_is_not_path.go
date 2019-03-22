package validators

import (
	"fmt"
	"os"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringSymlinkTargetIsNotPathError is a function that defines error message returned by StringSymlinkTargetIsNotPath validator.
// nolint: gochecknoglobals
var StringSymlinkTargetIsNotPathError = func(v *StringSymlinkTargetIsNotPath) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("symlink's '%s' target is an existing path", v.Field)
}

// StringSymlinkTargetIsNotPath is a validator object
type StringSymlinkTargetIsNotPath struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a symlink and it's target is an existing path.
func (v *StringSymlinkTargetIsNotPath) Validate(e *validator.Errors) {

	target, err := os.Readlink(v.Field)
	if err != nil {
		return
	}

	if _, err := os.Stat(target); os.IsNotExist(err) {
		return
	}

	e.Add(v.Name, StringSymlinkTargetIsNotPathError(v))
}

// SetField sets validator field.
func (v *StringSymlinkTargetIsNotPath) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringSymlinkTargetIsNotPath) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
