package buildin

import (
	"fmt"
	"net"

	"github.com/s3rj1k/validator"
)

// StringIsMACError is a function that defines error message returned by StringIsMAC validator.
// nolint: gochecknoglobals
var StringIsMACError = func(v *StringIsMAC) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be valid MAC address", v.Field)
}

// StringIsMAC is a validator object.
// Validate adds an error if the Field is not a MAC address.
type StringIsMAC struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a MAC address.
func (v *StringIsMAC) Validate(e *validator.Errors) {
	// using net ParseMAC
	_, err := net.ParseMAC(v.Field)
	if err == nil {
		return
	}

	e.Add(v.Name, StringIsMACError(v))
}

// SetField sets validator field.
func (v *StringIsMAC) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsMAC) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
