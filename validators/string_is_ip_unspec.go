package validators

import (
	"fmt"
	"net"

	"github.com/s3rj1k/validator"
)

// StringIsIPUnspecError is a function that defines error message returned by StringIsIPUnspec validator.
// nolint: gochecknoglobals
var StringIsIPUnspecError = func(v *StringIsIPUnspec) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be an unspecified address either IPv4 '0.0.0.0' or the IPv6 '::'", v.Field)
}

// StringIsIPUnspec is a validator object.
// Validate adds an error if the Field is not an unspecified address, either the IPv4 address "0.0.0.0" or the IPv6 address.
type StringIsIPUnspec struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not an unspecified address, either the IPv4 address "0.0.0.0" or the IPv6 address.
func (v *StringIsIPUnspec) Validate(e *validator.Errors) {

	ip := net.ParseIP(v.Field)

	if ip.IsUnspecified() {
		return
	}

	e.Add(v.Name, StringIsIPUnspecError(v))
}

// SetField sets validator field.
func (v *StringIsIPUnspec) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsIPUnspec) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
