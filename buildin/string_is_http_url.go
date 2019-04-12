package buildin

import (
	"fmt"
	"net/url"

	"github.com/s3rj1k/validator"
)

// StringIsHTTPURLError is a function that defines error message returned by StringIsHTTPURL validator.
// nolint: gochecknoglobals
var StringIsHTTPURLError = func(v *StringIsHTTPURL) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not a valid URL", v.Field)
}

// StringIsHTTPURL is a validator object.
// Validate adds an error if the Field is not a correctly formatted URL.
type StringIsHTTPURL struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a correctly formatted URL.
func (v *StringIsHTTPURL) Validate(e *validator.Errors) {

	var invalid = false

	if v.Field == "http://" || v.Field == "https://" {
		invalid = true
	}

	parsedURI, err := url.ParseRequestURI(v.Field)
	if err != nil {
		invalid = true
	}

	if parsedURI != nil && parsedURI.Scheme != "" && parsedURI.Scheme != "http" && parsedURI.Scheme != "https" {
		invalid = true
	}

	if !invalid {
		return
	}

	e.Add(v.Name, StringIsHTTPURLError(v))
}

// SetField sets validator field.
func (v *StringIsHTTPURL) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsHTTPURL) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
