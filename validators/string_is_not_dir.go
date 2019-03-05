package validators

import (
	"fmt"
	"os"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsNotDirError is a function that defines error message returned by StringIsNotDir validator.
// nolint: gochecknoglobals
var StringIsNotDirError = func(v *StringIsNotDir) string {
	return fmt.Sprintf("path '%s' is a dir", v.Field)
}

// StringIsNotDir is a validator object
type StringIsNotDir struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is a path to directory.
func (v *StringIsNotDir) Validate(e *validator.Errors) {

	fi, err := os.Stat(v.Field)
	if err != nil {
		return
	}

	if mode := fi.Mode(); !mode.IsDir() {
		return
	}

	e.Add(v.Name, StringIsNotDirError(v))
}

// SetField sets validator field.
func (v *StringIsNotDir) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNotDir) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
