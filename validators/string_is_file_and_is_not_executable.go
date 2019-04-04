package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsFileAndIsNotExecutableError is a function that defines error message returned by StringIsFileAndIsNotExecutable validator.
// nolint: gochecknoglobals
var StringIsFileAndIsNotExecutableError = func(v *StringIsFileAndIsNotExecutable) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' has execute bit set", v.Field)
}

// StringIsFileAndIsNotExecutable is a validator object.
// Validate adds an error if the Field is a file with execute bit set.
type StringIsFileAndIsNotExecutable struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a file with execute bit set.
func (v *StringIsFileAndIsNotExecutable) Validate(e *validator.Errors) {

	fi, err := os.Stat(v.Field)
	if err != nil {
		return
	}

	if fi.IsDir() || fi.Mode()&0111 != 0111 {
		return
	}

	e.Add(v.Name, StringIsFileAndIsNotExecutableError(v))
}

// SetField sets validator field.
func (v *StringIsFileAndIsNotExecutable) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsFileAndIsNotExecutable) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
