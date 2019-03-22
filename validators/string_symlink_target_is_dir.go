package validators

import (
	"fmt"
	"os"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringSymlinkTargetIsDirError is a function that defines error message returned by StringSymlinkTargetIsDir validator.
// nolint: gochecknoglobals
var StringSymlinkTargetIsDirError = func(v *StringSymlinkTargetIsDir) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("symlink's '%s' target is not a directory", v.Field)
}

// StringSymlinkTargetIsDir is a validator object
type StringSymlinkTargetIsDir struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a symlink and it's target is not a directory.
func (v *StringSymlinkTargetIsDir) Validate(e *validator.Errors) {

	target, err := os.Readlink(v.Field)
	if err != nil {
		return
	}

	if fi, err := os.Stat(target); !os.IsNotExist(err) {
		if mode := fi.Mode(); mode.IsDir() {
			return
		}
	}

	e.Add(v.Name, StringSymlinkTargetIsDirError(v))
}

// SetField sets validator field.
func (v *StringSymlinkTargetIsDir) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringSymlinkTargetIsDir) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
