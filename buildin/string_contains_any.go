package buildin

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringContainsAnyError is a function that defines error message returned by StringContainsAny validator.
// nolint: gochecknoglobals
var StringContainsAnyError = func(v *StringContainsAny) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' does not contain any substrings from %v", v.Field, v.ComparedField)
	}

	return fmt.Sprintf("'%s' does not contain any substrings from '%s'", v.Name, v.ComparedName)
}

// StringContainsAny is a validator object.
// Validate adds an error if the Field does not contain at least one substrings from ComparedField.
type StringContainsAny struct {
	Name          string
	Field         string
	ComparedName  string
	ComparedField []string
	Message       string
}

// Validate adds an error if the Field does not contain at least one substrings from ComparedField.
func (v *StringContainsAny) Validate(e *validator.Errors) {
	// if no required substrings - string is valid
	if v.ComparedField == nil || len(v.ComparedField) == 0 {
		return
	}

	for _, s := range v.ComparedField {
		if strings.Contains(v.Field, s) {
			return
		}
	}

	e.Add(v.Name, StringContainsAnyError(v))
}

// SetField sets validator field.
func (v *StringContainsAny) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringContainsAny) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
