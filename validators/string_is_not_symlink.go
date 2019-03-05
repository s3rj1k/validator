package validators

import (
	"fmt"
	"os"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsNotSymlinkError is a function that defines error message returned by StringIsNotSymlink validator.
// nolint: gochecknoglobals
var StringIsNotSymlinkError = func(v *StringIsNotSymlink) string {
	return fmt.Sprintf("path '%s' is a symlink", v.Field)
}

// StringIsNotSymlink is a validator object
type StringIsNotSymlink struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is a symlink.
func (v *StringIsNotSymlink) Validate(e *validator.Errors) {

	_, err := os.Readlink(v.Field)
	if err != nil {
		return
	}

	e.Add(v.Name, StringIsNotSymlinkError(v))
}

// SetField sets validator field.
func (v *StringIsNotSymlink) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNotSymlink) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
