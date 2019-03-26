package validators

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringHasSuffixNoneOfError is a function that defines error message returned by StringHasSuffixNoneOf validator.
// nolint: gochecknoglobals
var StringHasSuffixNoneOfError = func(v *StringHasSuffixNoneOf) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' has suffix from %v", v.Field, v.ComparedField)
	}

	return fmt.Sprintf("'%s' has suffix from contents of '%s'", v.Name, v.ComparedName)
}

// StringHasSuffixNoneOf is a validator object.
// Validate adds an error if the Field is Suffixed by at least one string from ComparedField.
type StringHasSuffixNoneOf struct {
	Name          string
	Field         string
	ComparedName  string
	ComparedField []string
	Message       string
}

// Validate adds an error if the Field is Suffixed by at least one string from ComparedField.
func (v *StringHasSuffixNoneOf) Validate(e *validator.Errors) {

	// if no excluding Suffixes - string is valid
	if v.ComparedField == nil || len(v.ComparedField) == 0 {
		return
	}

	var hasSuffix = false

	for _, s := range v.ComparedField {
		if strings.HasSuffix(v.Field, s) {
			hasSuffix = true
		}
	}

	if !hasSuffix {
		return
	}

	e.Add(v.Name, StringHasSuffixNoneOfError(v))
}

// SetField sets validator field.
func (v *StringHasSuffixNoneOf) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasSuffixNoneOf) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
