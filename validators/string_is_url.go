package validators

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsURLError is a function that defines error message returned by StringIsURL validator.
// nolint: gochecknoglobals
var StringIsURLError = func(v *StringIsURL) string {
	return fmt.Sprintf("%s is not a valid URL", v.Name)
}

// StringIsURL is a validator object
type StringIsURL struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not a correctly formatted URL.
func (v *StringIsURL) Validate(e *validator.Errors) {

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

	e.Add(v.Name, StringIsURLError(v))
}

// SetField sets validator field.
func (v *StringIsURL) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsURL) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
