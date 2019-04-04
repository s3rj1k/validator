package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsResolvableHostnameOrIPv6Error is a function that defines error message returned by StringIsResolvableHostnameOrIPv6 validator.
// nolint: gochecknoglobals
var StringIsResolvableHostnameOrIPv6Error = func(v *StringIsResolvableHostnameOrIPv6) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not a resolvable hostname and not an IPv6 address", v.Field)
}

// StringIsResolvableHostnameOrIPv6 is a validator object.
// Validate adds an error if the Field is not a resolvable hostname and not an IPv6 address.
type StringIsResolvableHostnameOrIPv6 struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a resolvable hostname and not an IPv6 address.
func (v *StringIsResolvableHostnameOrIPv6) Validate(e *validator.Errors) {

	if isResolvableHostname(v.Field) || isIPv6(v.Field) {
		return
	}

	e.Add(v.Name, StringIsResolvableHostnameOrIPv6Error(v))
}

// SetField sets validator field.
func (v *StringIsResolvableHostnameOrIPv6) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsResolvableHostnameOrIPv6) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
