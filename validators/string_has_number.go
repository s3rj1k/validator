package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringHasNumberError is a function that defines error message returned by StringHasNumber validator.
// nolint: gochecknoglobals
var StringHasNumberError = func(v *StringHasNumber) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' has no numbers", v.Field)
}

// StringHasNumber is a validator object.
// Validate adds an error if the Field has no numbers.
type StringHasNumber struct {
	Name    string
	Field   string
	Message string
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
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
