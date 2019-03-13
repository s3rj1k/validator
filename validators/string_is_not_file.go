package validators

import (
	"fmt"
	"os"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsNotFileError is a function that defines error message returned by StringIsNotFile validator.
// nolint: gochecknoglobals
var StringIsNotFileError = func(v *StringIsNotFile) string {
	return fmt.Sprintf("'%s' is a file", v.Name)
}

// StringIsNotFile is a validator object.
type StringIsNotFile struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is a file.
func (v *StringIsNotFile) Validate(e *validator.Errors) {

	fi, err := os.Stat(v.Field)
	if err != nil {
		return
	}

	if mode := fi.Mode(); mode.IsDir() {
		return
	}

	e.Add(v.Name, StringIsNotFileError(v))
}

// SetField sets validator field.
func (v *StringIsNotFile) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNotFile) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
