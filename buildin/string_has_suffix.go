package buildin

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringHasSuffixError is a function that defines error message returned by StringHasSuffix validator.
// nolint: gochecknoglobals
var StringHasSuffixError = func(v *StringHasSuffix) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' does not end with '%s'", v.Field, v.ComparedField)
	}

	return fmt.Sprintf("'%s' does not end with content of '%s'", v.Name, v.ComparedName)

}

// StringHasSuffix is a validator object.
// Validate adds an error if the Field is not suffixed with ComparedField.
type StringHasSuffix struct {
	Name          string
	Field         string
	ComparedName  string
	ComparedField string
	Message       string
}

// Validate adds an error if the Field is not suffixed with ComparedField.
func (v *StringHasSuffix) Validate(e *validator.Errors) {

	if strings.HasSuffix(v.Field, v.ComparedField) {
		return
	}

	e.Add(v.Name, StringHasSuffixError(v))
}

// SetField sets validator field.
func (v *StringHasSuffix) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasSuffix) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
