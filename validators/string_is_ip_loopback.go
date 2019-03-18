package validators

import (
	"fmt"
	"net"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsIPLoopbackError is a function that defines error message returned by StringIsIPLoopback validator.
// nolint: gochecknoglobals
var StringIsIPLoopbackError = func(v *StringIsIPLoopback) string {
	return fmt.Sprintf("'%s' must be a loopback address", v.Field)
}

// StringIsIPLoopback is a validator object.
type StringIsIPLoopback struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not a loopback address.
func (v *StringIsIPLoopback) Validate(e *validator.Errors) {

	ip := net.ParseIP(v.Field)

	if ip.IsLoopback() {
		return
	}

	e.Add(v.Name, StringIsIPLoopbackError(v))
}

// SetField sets validator field.
func (v *StringIsIPLoopback) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsIPLoopback) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
