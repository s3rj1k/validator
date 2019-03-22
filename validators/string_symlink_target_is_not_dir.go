package validators

import (
	"fmt"
	"os"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringSymlinkTargetIsNotDirError is a function that defines error message returned by StringSymlinkTargetIsNotDir validator.
// nolint: gochecknoglobals
var StringSymlinkTargetIsNotDirError = func(v *StringSymlinkTargetIsNotDir) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("symlink's '%s' target is a directory", v.Field)
}

// StringSymlinkTargetIsNotDir is a validator object
type StringSymlinkTargetIsNotDir struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a symlink and it's target is a directory.
func (v *StringSymlinkTargetIsNotDir) Validate(e *validator.Errors) {

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

	e.Add(v.Name, StringSymlinkTargetIsNotDirError(v))
}

// SetField sets validator field.
func (v *StringSymlinkTargetIsNotDir) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringSymlinkTargetIsNotDir) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
