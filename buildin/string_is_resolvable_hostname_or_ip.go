package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsResolvableHostnameOrIPError is a function that defines error message returned by StringIsResolvableHostnameOrIP validator.
// nolint: gochecknoglobals
var StringIsResolvableHostnameOrIPError = func(v *StringIsResolvableHostnameOrIP) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not a resolvable hostname and not an IP address", v.Field)
}

// StringIsResolvableHostnameOrIP is a validator object.
// Validate adds an error if the Field is not a resolvable hostname and not an IP address.
type StringIsResolvableHostnameOrIP struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a resolvable hostname and not an IP address.
func (v *StringIsResolvableHostnameOrIP) Validate(e *validator.Errors) {
	if isResolvableHostname(v.Field) || isIP(v.Field) {
		return
	}

	e.Add(v.Name, StringIsResolvableHostnameOrIPError(v))
}

// SetField sets validator field.
func (v *StringIsResolvableHostnameOrIP) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsResolvableHostnameOrIP) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
