package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsEmailError is a function that defines error message returned by StringIsEmail validator.
// nolint: gochecknoglobals
var StringIsEmailError = func(v *StringIsEmail) string {
	return fmt.Sprintf("'%s' does not match an email format", v.Field)
}

// StringIsEmail is a validator object.
type StringIsEmail struct {
	Name  string
	Field string
}

// Validate adds an error if the Field does not match email regexp. See Email const.
func (v *StringIsEmail) Validate(e *validator.Errors) {
	if rxEmail.Match([]byte(v.Field)) {
		return
	}

	e.Add(v.Name, StringIsEmailError(v))

}

// SetField sets validator field.
func (v *StringIsEmail) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsEmail) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
