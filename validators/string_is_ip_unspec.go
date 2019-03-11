package validators

import (
	"fmt"
	"net"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsIPUnspecError is a function that defines error message returned by StringIsIPUnspec validator.
// nolint: gochecknoglobals
var StringIsIPUnspecError = func(v *StringIsIPUnspec) string {
	return fmt.Sprintf("'%s' must be an unspecified address", v.Name)
}

// StringIsIPUnspec is a validator object.
type StringIsIPUnspec struct {
	Name  string
	Field string
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
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
