package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsNotResolvableHostnameNorIPv4Error is a function that defines error message returned by StringIsNotResolvableHostnameNorIPv4 validator.
// nolint: gochecknoglobals
var StringIsNotResolvableHostnameNorIPv4Error = func(v *StringIsNotResolvableHostnameNorIPv4) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is a resolvable hostname or an IPv4 address", v.Field)
}

// StringIsNotResolvableHostnameNorIPv4 is a validator object.
// Validate adds an error if the Field is a resolvable hostname or an IPv4 address.
type StringIsNotResolvableHostnameNorIPv4 struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a resolvable hostname or an IPv4 address.
func (v *StringIsNotResolvableHostnameNorIPv4) Validate(e *validator.Errors) {

	if !isResolvableHostname(v.Field) && !isIPv4(v.Field) {
		return
	}

	e.Add(v.Name, StringIsNotResolvableHostnameNorIPv4Error(v))
}

// SetField sets validator field.
func (v *StringIsNotResolvableHostnameNorIPv4) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNotResolvableHostnameNorIPv4) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
