package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringIsFileAndIsExecutableError is a function that defines error message returned by StringIsFileAndIsExecutable validator.
// nolint: gochecknoglobals
var StringIsFileAndIsExecutableError = func(v *StringIsFileAndIsExecutable) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' does not have execute bit set", v.Field)
}

// StringIsFileAndIsExecutable is a validator object.
// Validate adds an error if the Field is a file without execute bit set.
type StringIsFileAndIsExecutable struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a file without execute bit set.
func (v *StringIsFileAndIsExecutable) Validate(e *validator.Errors) {

	fi, err := os.Stat(v.Field)
	if err != nil {
		return
	}

	if fi.IsDir() {
		return
	}

	if fi.Mode()&0111 == 0111 {
		return
	}

	e.Add(v.Name, StringIsFileAndIsExecutableError(v))
}

// SetField sets validator field.
func (v *StringIsFileAndIsExecutable) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsFileAndIsExecutable) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
