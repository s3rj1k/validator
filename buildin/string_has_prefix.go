package buildin

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringHasPrefixError is a function that defines error message returned by StringHasPrefix validator.
// nolint: gochecknoglobals
var StringHasPrefixError = func(v *StringHasPrefix) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' does not start with '%s'", v.Field, v.ComparedField)
	}

	return fmt.Sprintf("'%s' does not start with content of '%s'", v.Name, v.ComparedName)
}

// StringHasPrefix is a validator object.
// Validate adds an error if the Field is not prefixed with ComparedField.
type StringHasPrefix struct {
	Name          string
	Field         string
	ComparedName  string
	ComparedField string
	Message       string
}

// Validate adds an error if the Field is not prefixed with ComparedField.
func (v *StringHasPrefix) Validate(e *validator.Errors) {
	if strings.HasPrefix(v.Field, v.ComparedField) {
		return
	}

	e.Add(v.Name, StringHasPrefixError(v))
}

// SetField sets validator field.
func (v *StringHasPrefix) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasPrefix) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
