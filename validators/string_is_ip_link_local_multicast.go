package validators

import (
	"fmt"
	"net"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsIPLinkLocalMulticastError is a function that defines error message returned by StringIsIPLinkLocalMulticast validator.
// nolint: gochecknoglobals
var StringIsIPLinkLocalMulticastError = func(v *StringIsIPLinkLocalMulticast) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be a link-local multicast address", v.Field)
}

// StringIsIPLinkLocalMulticast is a validator object.
type StringIsIPLinkLocalMulticast struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a link-local multicast address.
func (v *StringIsIPLinkLocalMulticast) Validate(e *validator.Errors) {

	ip := net.ParseIP(v.Field)

	if ip.IsLinkLocalMulticast() {
		return
	}

	e.Add(v.Name, StringIsIPLinkLocalMulticastError(v))
}

// SetField sets validator field.
func (v *StringIsIPLinkLocalMulticast) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsIPLinkLocalMulticast) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
