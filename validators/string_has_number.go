package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringHasNumberError is a function that defines error message returned by StringHasNumber validator.
// nolint: gochecknoglobals
var StringHasNumberError = func(v *StringHasNumber) string {
	return fmt.Sprintf("'%s' has no numbers", v.Name)
}

// StringHasNumber is a validator object.
type StringHasNumber struct {
	Name  string
	Field string
}

// Validate adds an error if the Field has no numbers.
func (v *StringHasNumber) Validate(e *validator.Errors) {

	if rxHasNumber.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringHasNumberError(v))
}

// SetField sets validator field.
func (v *StringHasNumber) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasNumber) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
