package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsNotResolvableHostnameError is a function that defines error message returned by StringIsNotResolvableHostname validator.
// nolint: gochecknoglobals
var StringIsNotResolvableHostnameError = func(v *StringIsNotResolvableHostname) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is a resolvable hostname", v.Field)
}

// StringIsNotResolvableHostname is a validator object.
// Validate adds an error if the Field is a resolvable hostname.
type StringIsNotResolvableHostname struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a resolvable hostname.
func (v *StringIsNotResolvableHostname) Validate(e *validator.Errors) {

	if !isResolvableHostname(v.Field) {
		return
	}

	e.Add(v.Name, StringIsNotResolvableHostnameError(v))
}

// SetField sets validator field.
func (v *StringIsNotResolvableHostname) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNotResolvableHostname) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
