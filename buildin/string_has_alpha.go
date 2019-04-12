package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringHasAlphaError is a function that defines error message returned by StringHasAlpha validator.
// nolint: gochecknoglobals
var StringHasAlphaError = func(v *StringHasAlpha) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' has no letters", v.Field)
}

// StringHasAlpha is a validator object.
// Validate adds an error if the Field has no letters.
type StringHasAlpha struct {
	Name    string
	Field   string
	Message string
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
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
