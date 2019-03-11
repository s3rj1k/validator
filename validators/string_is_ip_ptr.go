package validators

import (
	"fmt"
	"net"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsIPptrError is a function that defines error message returned by StringIsIPptr validator.
// nolint: gochecknoglobals
var StringIsIPptrError = func(v *StringIsIPptr) string {
	return fmt.Sprintf("'%s' must be an IP address with available PTR record", v.Name)
}

// StringIsIPptr is a validator object.
type StringIsIPptr struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is an IP address that does not have PTR record.
func (v *StringIsIPptr) Validate(e *validator.Errors) {

	names, err := net.LookupAddr(v.Field)
	if err == nil || names != nil || len(names) != 0 {
		return
	}

	e.Add(v.Name, StringIsIPptrError(v))
}

// SetField sets validator field.
func (v *StringIsIPptr) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsIPptr) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
