package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
	"golang.org/x/sys/unix"
)

// StringIsPathAndIsWritableError is a function that defines error message returned by StringIsPathAndIsWritable validator.
// nolint: gochecknoglobals
var StringIsPathAndIsWritableError = func(v *StringIsPathAndIsWritable) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' path is not writable", v.Field)
}

// StringIsPathAndIsWritable is a validator object.
// Validate adds an error if the Field is a path and is not writable.
type StringIsPathAndIsWritable struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a path and is not writable.
func (v *StringIsPathAndIsWritable) Validate(e *validator.Errors) {

	if err := unix.Access(v.Field, unix.W_OK); err == nil {
		return
	}

	e.Add(v.Name, StringIsPathAndIsWritableError(v))
}

// SetField sets validator field.
func (v *StringIsPathAndIsWritable) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPathAndIsWritable) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
