package validators

import (
	"fmt"
	"net"

	"github.com/s3rj1k/validator"
)

// StringIsCIDRv6Error is a function that defines error message returned by StringIsCIDRv6 validator.
// nolint: gochecknoglobals
var StringIsCIDRv6Error = func(v *StringIsCIDRv6) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be CIDR notation of IPv6 address", v.Field)
}

// StringIsCIDRv6 is a validator object.
// Validate adds an error if the Field is not a CIDR notation of IPv6 address.
type StringIsCIDRv6 struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a CIDR notation of IPv6 address.
func (v *StringIsCIDRv6) Validate(e *validator.Errors) {

	ip, _, err := net.ParseCIDR(v.Field)
	if err == nil && ip.To4() == nil {
		return
	}

	e.Add(v.Name, StringIsCIDRv6Error(v))
}

// SetField sets validator field.
func (v *StringIsCIDRv6) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsCIDRv6) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
