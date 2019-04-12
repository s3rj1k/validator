package buildin

import (
	"fmt"
	"net"

	"github.com/s3rj1k/validator"
)

// StringIsIPMulticastError is a function that defines error message returned by StringIsIPMulticast validator.
// nolint: gochecknoglobals
var StringIsIPMulticastError = func(v *StringIsIPMulticast) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be a multicast address", v.Field)
}

// StringIsIPMulticast is a validator object.
// Validate adds an error if the Field is not a multicast address.
type StringIsIPMulticast struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a multicast address.
func (v *StringIsIPMulticast) Validate(e *validator.Errors) {

	ip := net.ParseIP(v.Field)

	if ip.IsMulticast() {
		return
	}

	e.Add(v.Name, StringIsIPMulticastError(v))
}

// SetField sets validator field.
func (v *StringIsIPMulticast) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsIPMulticast) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
