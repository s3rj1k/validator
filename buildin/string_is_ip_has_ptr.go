package buildin

import (
	"fmt"
	"net"

	"github.com/s3rj1k/validator"
)

// StringIsIPHasPTRError is a function that defines error message returned by StringIsIPHasPTR validator.
// nolint: gochecknoglobals
var StringIsIPHasPTRError = func(v *StringIsIPHasPTR) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be an IP address with resolvable PTR record", v.Field)
}

// StringIsIPHasPTR is a validator object.
// Validate adds an error if the Field is an IP address that does not have PTR record.
type StringIsIPHasPTR struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is an IP address that does not have PTR record.
func (v *StringIsIPHasPTR) Validate(e *validator.Errors) {

	names, err := net.LookupAddr(v.Field)
	if err == nil || names != nil || len(names) > 0 {
		return
	}

	e.Add(v.Name, StringIsIPHasPTRError(v))
}

// SetField sets validator field.
func (v *StringIsIPHasPTR) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsIPHasPTR) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
