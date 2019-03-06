package validators

import (
	"fmt"
	"os"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsSymlinkError is a function that defines error message returned by StringIsSymlink validator.
// nolint: gochecknoglobals
var StringIsSymlinkError = func(v *StringIsSymlink) string {
	return fmt.Sprintf("path '%s' is not a symlink", v.Field)
}

// StringIsSymlink is a validator object
type StringIsSymlink struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not a symlink.
func (v *StringIsSymlink) Validate(e *validator.Errors) {

	_, err := os.Readlink(v.Field)
	if err == nil {
		return
	}

	e.Add(v.Name, StringIsSymlinkError(v))
}

// SetField sets validator field.
func (v *StringIsSymlink) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsSymlink) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
