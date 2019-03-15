package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringHasNoWhitespaceError is a function that defines error message returned by StringHasNoWhitespace validator.
// nolint: gochecknoglobals
var StringHasNoWhitespaceError = func(v *StringHasNoWhitespace) string {
	return fmt.Sprintf("'%s' has whitespace", v.Field)
}

// StringHasNoWhitespace is a validator object.
type StringHasNoWhitespace struct {
	Name  string
	Field string
}

// Validate adds an error if the Field has whitespace.
func (v *StringHasNoWhitespace) Validate(e *validator.Errors) {

	if !strings.Contains(v.Field, " ") {
		return
	}

	e.Add(v.Name, StringHasNoWhitespaceError(v))
}

// SetField sets validator field.
func (v *StringHasNoWhitespace) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasNoWhitespace) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
