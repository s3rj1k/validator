package buildin

import (
	"fmt"
	"net"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsIPv6Error is a function that defines error message returned by StringIsIPv6 validator.
// nolint: gochecknoglobals
var StringIsIPv6Error = func(v *StringIsIPv6) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be IPv6", v.Field)
}

// StringIsIPv6 is a validator object.
// Validate adds an error if the Field is not a valid IPv6 address.
type StringIsIPv6 struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a valid IPv6 address.
func (v *StringIsIPv6) Validate(e *validator.Errors) {

	if isIPv6(v.Field) {
		return
	}

	e.Add(v.Name, StringIsIPv6Error(v))
}

// SetField sets validator field.
func (v *StringIsIPv6) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsIPv6) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}

// checks if s is a valid IPv6
func isIPv6(s string) bool {

	if len(s) == 0 {
		return false
	}

	ip := net.ParseIP(s)
	if ip.To16() == nil || len(ip) != net.IPv6len || !strings.Contains(s, ":") {
		return false
	}

	return true
}
