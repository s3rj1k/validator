package buildin

import (
	"fmt"
	"strconv"

	"github.com/s3rj1k/validator"
)

// StringIsPortError is a function that defines error message returned by StringIsPort validator.
// nolint: gochecknoglobals
var StringIsPortError = func(v *StringIsPort) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not a valid port", v.Field)
}

// StringIsPort is a validator object.
// Validate adds an error if the Field does not represent a valid port.
type StringIsPort struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field does not represent a valid port.
func (v *StringIsPort) Validate(e *validator.Errors) {

	if i, err := strconv.Atoi(v.Field); err == nil && i > 0 && i < 65536 {
		return
	}

	e.Add(v.Name, StringIsPortError(v))
}

// SetField sets validator field.
func (v *StringIsPort) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsPort) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
