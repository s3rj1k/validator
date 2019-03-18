package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringMatchRegexError is a function that defines error message returned by StringMatchRegex validator.
// nolint: gochecknoglobals
var StringMatchRegexError = func(v *StringMatchRegex) string {
	return fmt.Sprintf("'%s' does not match regex '%s'", v.Field, v.Regex)
}

// StringMatchRegex is a validator object.
type StringMatchRegex struct {
	Name  string
	Field string
	Regex string
}

// Validate adds an error if the Field does not match regular expression Regex.
func (v *StringMatchRegex) Validate(e *validator.Errors) {
	r := regexp.MustCompile(v.Regex)
	if r.Match([]byte(v.Field)) {
		return
	}

	e.Add(v.Name, StringMatchRegexError(v))
}

// SetField sets validator field.
func (v *StringMatchRegex) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringMatchRegex) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
