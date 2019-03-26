package validators

import (
	"fmt"
	"os"

	"github.com/s3rj1k/validator"
)

// StringFileIsExecutableError is a function that defines error message returned by StringFileIsExecutable validator.
// nolint: gochecknoglobals
var StringFileIsExecutableError = func(v *StringFileIsExecutable) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' does not have execute bit set", v.Field)
}

// StringFileIsExecutable is a validator object.
// Validate adds an error if the Field is a file without execute bit set.
type StringFileIsExecutable struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a file without execute bit set.
func (v *StringFileIsExecutable) Validate(e *validator.Errors) {

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

	e.Add(v.Name, StringFileIsExecutableError(v))
}

// SetField sets validator field.
func (v *StringFileIsExecutable) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringFileIsExecutable) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
