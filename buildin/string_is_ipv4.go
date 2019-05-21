package buildin

import (
	"fmt"
	"net"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsIPv4Error is a function that defines error message returned by StringIsIPv4 validator.
// nolint: gochecknoglobals
var StringIsIPv4Error = func(v *StringIsIPv4) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be IPv4", v.Field)
}

// StringIsIPv4 is a validator object.
// Validate adds an error if the Field is not a valid IPv4 address.
type StringIsIPv4 struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a valid IPv4 address.
func (v *StringIsIPv4) Validate(e *validator.Errors) {
	if isIPv4(v.Field) {
		return
	}

	e.Add(v.Name, StringIsIPv4Error(v))
}

// SetField sets validator field.
func (v *StringIsIPv4) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsIPv4) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}

// checks if s is a valid IPv4
func isIPv4(s string) bool {
	if len(s) == 0 {
		return false
	}

	ip := net.ParseIP(s)
	if ip.To4() == nil || !strings.Contains(s, ".") {
		return false
	}

	return true
}
