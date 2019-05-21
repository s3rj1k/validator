package buildin

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsNotSymlinkError is a function that defines error message returned by StringIsNotSymlink validator.
// nolint: gochecknoglobals
var StringIsNotSymlinkError = func(v *StringIsNotSymlink) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("path '%s' is a symlink", v.Field)
}

// StringIsNotSymlink is a validator object.
// Validate adds an error if the Field is a symlink.
type StringIsNotSymlink struct {
	Name    string
	Field   string
	Message string
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
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
