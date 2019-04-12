package buildin

import (
	"fmt"
	"net"

	"github.com/s3rj1k/validator"
)

// StringIsIPLinkLocalUnicastError is a function that defines error message returned by StringIsIPLinkLocalUnicast validator.
// nolint: gochecknoglobals
var StringIsIPLinkLocalUnicastError = func(v *StringIsIPLinkLocalUnicast) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be a link-local unicast address", v.Field)
}

// StringIsIPLinkLocalUnicast is a validator object.
// Validate adds an error if the Field is not a link-local unicast address.
type StringIsIPLinkLocalUnicast struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a link-local unicast address.
func (v *StringIsIPLinkLocalUnicast) Validate(e *validator.Errors) {

	ip := net.ParseIP(v.Field)

	if ip.IsLinkLocalUnicast() {
		return
	}

	e.Add(v.Name, StringIsIPLinkLocalUnicastError(v))
}

// SetField sets validator field.
func (v *StringIsIPLinkLocalUnicast) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsIPLinkLocalUnicast) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
