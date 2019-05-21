package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsResolvableHostnameOrIPv4Error is a function that defines error message returned by StringIsResolvableHostnameOrIPv4 validator.
// nolint: gochecknoglobals
var StringIsResolvableHostnameOrIPv4Error = func(v *StringIsResolvableHostnameOrIPv4) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not a resolvable hostname and not an IPv4 address", v.Field)
}

// StringIsResolvableHostnameOrIPv4 is a validator object.
// Validate adds an error if the Field is not a resolvable hostname and not an IPv4 address.
type StringIsResolvableHostnameOrIPv4 struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a resolvable hostname and not an IPv4 address.
func (v *StringIsResolvableHostnameOrIPv4) Validate(e *validator.Errors) {
	if isResolvableHostname(v.Field) || isIPv4(v.Field) {
		return
	}

	e.Add(v.Name, StringIsResolvableHostnameOrIPv4Error(v))
}

// SetField sets validator field.
func (v *StringIsResolvableHostnameOrIPv4) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsResolvableHostnameOrIPv4) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
