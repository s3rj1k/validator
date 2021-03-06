package buildin

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringContainsNoneOfError is a function that defines error message returned by StringContainsNoneOf validator.
// nolint: gochecknoglobals
var StringContainsNoneOfError = func(v *StringContainsNoneOf) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' contains substring from %v", v.Field, v.ComparedField)
	}

	return fmt.Sprintf("'%s' contains substring from '%s'", v.Name, v.ComparedName)
}

// StringContainsNoneOf is a validator object.
// Validate adds an error if the Field contains at least one substrings from ComparedField.
type StringContainsNoneOf struct {
	Name          string
	Field         string
	ComparedName  string
	ComparedField []string
	Message       string
}

// Validate adds an error if the Field contains at least one substrings from ComparedField.
func (v *StringContainsNoneOf) Validate(e *validator.Errors) {
	// if no excluding substrings - string is valid
	if v.ComparedField == nil || len(v.ComparedField) == 0 {
		return
	}

	var hasSubstring = false

	for _, s := range v.ComparedField {
		if strings.Contains(v.Field, s) {
			hasSubstring = true
		}
	}

	if !hasSubstring {
		return
	}

	e.Add(v.Name, StringContainsNoneOfError(v))
}

// SetField sets validator field.
func (v *StringContainsNoneOf) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringContainsNoneOf) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
