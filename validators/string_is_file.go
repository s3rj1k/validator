package validators

import (
	"fmt"
	"os"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsFileError is a function that defines error message returned by StringIsFile validator.
// nolint: gochecknoglobals
var StringIsFileError = func(v *StringIsFile) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not a file", v.Field)
}

// StringIsFile is a validator object.
type StringIsFile struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a file.
func (v *StringIsFile) Validate(e *validator.Errors) {

	if fi, err := os.Stat(v.Field); !os.IsNotExist(err) {
		if !fi.IsDir() {
			return
		}
	}

	e.Add(v.Name, StringIsFileError(v))
}

// SetField sets validator field.
func (v *StringIsFile) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsFile) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
