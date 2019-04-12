package buildin

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringHasSuffixAnyError is a function that defines error message returned by StringHasSuffixAny validator.
// nolint: gochecknoglobals
var StringHasSuffixAnyError = func(v *StringHasSuffixAny) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' has no suffix from %v", v.Field, v.ComparedField)
	}

	return fmt.Sprintf("'%s' has no suffix from contents of '%s'", v.Name, v.ComparedName)
}

// StringHasSuffixAny is a validator object.
// Validate adds an error if the Field is not Suffixed by at least one string from ComparedField.
type StringHasSuffixAny struct {
	Name          string
	Field         string
	ComparedName  string
	ComparedField []string
	Message       string
}

// Validate adds an error if the Field is not Suffixed by at least one string from ComparedField.
func (v *StringHasSuffixAny) Validate(e *validator.Errors) {

	// if no required suffixes - string is valid
	if v.ComparedField == nil || len(v.ComparedField) == 0 {
		return
	}

	for _, s := range v.ComparedField {
		if strings.HasSuffix(v.Field, s) {
			return
		}
	}

	e.Add(v.Name, StringHasSuffixAnyError(v))
}

// SetField sets validator field.
func (v *StringHasSuffixAny) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasSuffixAny) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
