package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsDirError is a function that defines error message returned by StringIsDir validator.
// nolint: gochecknoglobals
var StringIsDirError = func(v *StringIsDir) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("path '%s' is not a dir", v.Field)
}

// StringIsDir is a validator object.
// Validate adds an error if the Field is not a path to directory.
type StringIsDir struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a path to directory.
func (v *StringIsDir) Validate(e *validator.Errors) {
	if fi, err := os.Stat(v.Field); !os.IsNotExist(err) {
		if mode := fi.Mode(); mode.IsDir() {
			return
		}
	}

	e.Add(v.Name, StringIsDirError(v))
}

// SetField sets validator field.
func (v *StringIsDir) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsDir) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
