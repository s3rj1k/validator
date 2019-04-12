package buildin

import (
	"fmt"
	"net"

	"github.com/s3rj1k/validator"
)

// StringIsResolvableHostnameError is a function that defines error message returned by StringIsResolvableHostname validator.
// nolint: gochecknoglobals
var StringIsResolvableHostnameError = func(v *StringIsResolvableHostname) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not a resolvable hostname", v.Field)
}

// StringIsResolvableHostname is a validator object.
// Validate adds an error if the Field is not a resolvable hostname.
type StringIsResolvableHostname struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a resolvable hostname.
func (v *StringIsResolvableHostname) Validate(e *validator.Errors) {

	if isResolvableHostname(v.Field) {
		return
	}

	e.Add(v.Name, StringIsResolvableHostnameError(v))
}

// SetField sets validator field.
func (v *StringIsResolvableHostname) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsResolvableHostname) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}

// isResolvableHostname if s is a resolvable host NAME. IP addresses return false.
func isResolvableHostname(s string) bool {

	// resolvable host is OK
	addrs, err := net.LookupHost(s)
	if err != nil || len(addrs) == 0 {
		return false
	}

	// IP addr is BAD
	if ip := net.ParseIP(s); ip != nil {
		return false
	}

	return true
}
