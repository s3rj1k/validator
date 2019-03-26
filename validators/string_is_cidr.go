package validators

import (
	"fmt"
	"net"

	"github.com/s3rj1k/validator"
)

// StringIsCIDRError is a function that defines error message returned by StringIsCIDR validator.
// nolint: gochecknoglobals
var StringIsCIDRError = func(v *StringIsCIDR) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be a CIDR notation address", v.Field)
}

// StringIsCIDR is a validator object.
// Validate adds an error if the Field is not a valid CIDR notation address.
type StringIsCIDR struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a valid CIDR notation address.
func (v *StringIsCIDR) Validate(e *validator.Errors) {

	_, _, err := net.ParseCIDR(v.Field)
	if err == nil {
		return
	}

	e.Add(v.Name, StringIsCIDRError(v))
}

// SetField sets validator field.
func (v *StringIsCIDR) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsCIDR) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
