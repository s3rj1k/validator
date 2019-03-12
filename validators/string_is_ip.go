package validators

import (
	"fmt"
	"net"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsIPError is a function that defines error message returned by StringIsIP validator.
// nolint: gochecknoglobals
var StringIsIPError = func(v *StringIsIP) string {
	return fmt.Sprintf("%s must be either IP version 4 or 6", v.Name)
}

// StringIsIP is a validator object.
type StringIsIP struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is a valid IP address version 4 or 6.
func (v *StringIsIP) Validate(e *validator.Errors) {

	if isIP(v.Field) {
		return
	}

	e.Add(v.Name, StringIsIPError(v))
}

// SetField sets validator field.
func (v *StringIsIP) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsIP) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}

// checks if s is a valid IP
func isIP(s string) bool {

	if len(s) == 0 {
		return false
	}

	ip := net.ParseIP(s)
	return ip != nil
}
