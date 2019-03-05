package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringHasNoSuffixError is a function that defines error message returned by StringHasNoSuffix validator.
// nolint: gochecknoglobals
var StringHasNoSuffixError = func(v *StringHasNoSuffix) string {
	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' ends with '%s'", v.Field, v.ComparedField)
	}

	return fmt.Sprintf("'%s' ends with content of '%s'", v.Name, v.ComparedName)
}

// StringHasNoSuffix is a validator object.
type StringHasNoSuffix struct {
	Name          string
	Field         string
	ComparedName  string
	ComparedField string
}

// Validate adds an error if the Field is suffixed with ComparedField.
func (v *StringHasNoSuffix) Validate(e *validator.Errors) {

	if !strings.HasSuffix(v.Field, v.ComparedField) {
		return
	}

	e.Add(v.Name, StringHasNoSuffixError(v))

}

// SetField sets validator field.
func (v *StringHasNoSuffix) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasNoSuffix) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
