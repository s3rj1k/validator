package validators

import (
	"fmt"
	"net"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsIPIfaceLocalMulticastError is a function that defines error message returned by StringIsIPIfaceLocalMulticast validator.
// nolint: gochecknoglobals
var StringIsIPIfaceLocalMulticastError = func(v *StringIsIPIfaceLocalMulticast) string {
	return fmt.Sprintf("'%s' must be an interface-local multicast address", v.Field)
}

// StringIsIPIfaceLocalMulticast is a validator object.
type StringIsIPIfaceLocalMulticast struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not an interface-local multicast address.
func (v *StringIsIPIfaceLocalMulticast) Validate(e *validator.Errors) {

	ip := net.ParseIP(v.Field)

	if ip.IsInterfaceLocalMulticast() {
		return
	}

	e.Add(v.Name, StringIsIPIfaceLocalMulticastError(v))
}

// SetField sets validator field.
func (v *StringIsIPIfaceLocalMulticast) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsIPIfaceLocalMulticast) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
