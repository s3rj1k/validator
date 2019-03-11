package validators

import (
	"fmt"
	"net"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsIPLinkLocalUnicastError is a function that defines error message returned by StringIsIPLinkLocalUnicast validator.
// nolint: gochecknoglobals
var StringIsIPLinkLocalUnicastError = func(v *StringIsIPLinkLocalUnicast) string {
	return fmt.Sprintf("'%s' must be a link-local unicast address", v.Name)
}

// StringIsIPLinkLocalUnicast is a validator object.
type StringIsIPLinkLocalUnicast struct {
	Name  string
	Field string
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
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
