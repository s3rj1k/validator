package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringHasAlphaError is a function that defines error message returned by StringHasAlpha validator.
// nolint: gochecknoglobals
var StringHasAlphaError = func(v *StringHasAlpha) string {
	return fmt.Sprintf("'%s' has no letters", v.Field)
}

// StringHasAlpha is a validator object.
type StringHasAlpha struct {
	Name  string
	Field string
}

// Validate adds an error if the Field has no letters.
func (v *StringHasAlpha) Validate(e *validator.Errors) {

	if rxHasAlpha.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringHasAlphaError(v))
}

// SetField sets validator field.
func (v *StringHasAlpha) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasAlpha) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
