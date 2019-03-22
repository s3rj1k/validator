package validators

import (
	"fmt"
	"net"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsIPGlobalUnicastError is a function that defines error message returned by StringIsIPGlobalUnicast validator.
// nolint: gochecknoglobals
var StringIsIPGlobalUnicastError = func(v *StringIsIPGlobalUnicast) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be Global Unicast IP", v.Field)
}

// StringIsIPGlobalUnicast is a validator object.
type StringIsIPGlobalUnicast struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a valid Global Unicast IP address.
// Error is not added if ip is in IPv4 private address space or
// local IPv6 unicast address space.
func (v *StringIsIPGlobalUnicast) Validate(e *validator.Errors) {

	ip := net.ParseIP(v.Field)

	if ip.IsGlobalUnicast() {
		return
	}

	e.Add(v.Name, StringIsIPGlobalUnicastError(v))
}

// SetField sets validator field.
func (v *StringIsIPGlobalUnicast) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsIPGlobalUnicast) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
