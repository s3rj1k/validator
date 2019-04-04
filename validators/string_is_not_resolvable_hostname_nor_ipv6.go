package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsNotResolvableHostnameNorIPv6Error is a function that defines error message returned by StringIsNotResolvableHostnameNorIPv6 validator.
// nolint: gochecknoglobals
var StringIsNotResolvableHostnameNorIPv6Error = func(v *StringIsNotResolvableHostnameNorIPv6) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is a resolvable hostname or an IPv6 address", v.Field)
}

// StringIsNotResolvableHostnameNorIPv6 is a validator object.
// Validate adds an error if the Field is a resolvable hostname or an IPv6 address.
type StringIsNotResolvableHostnameNorIPv6 struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a resolvable hostname or an IPv6 address.
func (v *StringIsNotResolvableHostnameNorIPv6) Validate(e *validator.Errors) {

	if !isResolvableHostname(v.Field) && !isIPv6(v.Field) {
		return
	}

	e.Add(v.Name, StringIsNotResolvableHostnameNorIPv6Error(v))
}

// SetField sets validator field.
func (v *StringIsNotResolvableHostnameNorIPv6) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNotResolvableHostnameNorIPv6) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
