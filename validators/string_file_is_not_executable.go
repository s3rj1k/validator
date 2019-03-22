package validators

import (
	"fmt"
	"os"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringFileIsNotExecutableError is a function that defines error message returned by StringFileIsNotExecutable validator.
// nolint: gochecknoglobals
var StringFileIsNotExecutableError = func(v *StringFileIsNotExecutable) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' has execute bit set", v.Field)
}

// StringFileIsNotExecutable is a validator object.
type StringFileIsNotExecutable struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a file with execute bit set.
func (v *StringFileIsNotExecutable) Validate(e *validator.Errors) {

	fi, err := os.Stat(v.Field)
	if err != nil {
		return
	}

	if fi.IsDir() || fi.Mode()&0111 != 0111 {
		return
	}

	e.Add(v.Name, StringFileIsNotExecutableError(v))
}

// SetField sets validator field.
func (v *StringFileIsNotExecutable) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringFileIsNotExecutable) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
