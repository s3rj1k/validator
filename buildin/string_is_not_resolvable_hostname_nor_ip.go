package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsNotResolvableHostnameNorIPError is a function that defines error message returned by StringIsNotResolvableHostnameNorIP validator.
// nolint: gochecknoglobals
var StringIsNotResolvableHostnameNorIPError = func(v *StringIsNotResolvableHostnameNorIP) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is a resolvable hostname or an IP address", v.Field)
}

// StringIsNotResolvableHostnameNorIP is a validator object.
// Validate adds an error if the Field is a resolvable hostname or an IP address.
type StringIsNotResolvableHostnameNorIP struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a resolvable hostname or an IP address.
func (v *StringIsNotResolvableHostnameNorIP) Validate(e *validator.Errors) {
	if !isResolvableHostname(v.Field) && !isIP(v.Field) {
		return
	}

	e.Add(v.Name, StringIsNotResolvableHostnameNorIPError(v))
}

// SetField sets validator field.
func (v *StringIsNotResolvableHostnameNorIP) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNotResolvableHostnameNorIP) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
