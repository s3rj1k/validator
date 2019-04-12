package buildin

import (
	"fmt"
	"net"

	"github.com/s3rj1k/validator"
)

// StringIsCIDRv4Error is a function that defines error message returned by StringIsCIDRv4 validator.
// nolint: gochecknoglobals
var StringIsCIDRv4Error = func(v *StringIsCIDRv4) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be CIDR notation of IPv4 address", v.Field)
}

// StringIsCIDRv4 is a validator object.
// Validate adds an error if the Field is not a valid CIDR notation of IPv4 address.
type StringIsCIDRv4 struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a valid CIDR notation of IPv4 address.
func (v *StringIsCIDRv4) Validate(e *validator.Errors) {

	ip, _, err := net.ParseCIDR(v.Field)
	if err == nil && ip.To4() != nil {
		return
	}

	e.Add(v.Name, StringIsCIDRv4Error(v))
}

// SetField sets validator field.
func (v *StringIsCIDRv4) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsCIDRv4) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
