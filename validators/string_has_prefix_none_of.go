package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringHasPrefixNoneOfError is a function that defines error message returned by StringHasPrefixNoneOf validator.
// nolint: gochecknoglobals
var StringHasPrefixNoneOfError = func(v *StringHasPrefixNoneOf) string {
	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' has prefix from '%v'", v.Name, v.ComparedField)
	}

	return fmt.Sprintf("'%s' has prefix from '%s'", v.Name, v.ComparedName)
}

// StringHasPrefixNoneOf is a validator object.
type StringHasPrefixNoneOf struct {
	Name          string
	Field         string
	ComparedName  string
	ComparedField []string
}

// Validate adds an error if the Field is prefixed by at least one string from ComparedField.
func (v *StringHasPrefixNoneOf) Validate(e *validator.Errors) {

	// if no excluding prefixes - string is valid
	if v.ComparedField == nil || len(v.ComparedField) == 0 {
		return
	}

	var hasPrefix = false

	for _, s := range v.ComparedField {
		if strings.HasPrefix(v.Field, s) {
			hasPrefix = true
		}
	}

	if !hasPrefix {
		return
	}

	e.Add(v.Name, StringHasPrefixNoneOfError(v))
}

// SetField sets validator field.
func (v *StringHasPrefixNoneOf) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasPrefixNoneOf) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
