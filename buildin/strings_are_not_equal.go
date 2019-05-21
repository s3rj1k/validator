package buildin

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringsAreNotEqualError is a function that defines error message returned by StringsAreNotEqual validator.
// nolint: gochecknoglobals
var StringsAreNotEqualError = func(v *StringsAreNotEqual) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	var caseName string

	if v.CaseInsensitive {
		caseName = "iequal"
	} else {
		caseName = "equal"
	}

	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' is %s to '%s'", v.Field, caseName, v.ComparedField)
	}

	return fmt.Sprintf("'%s' is %s to '%s'", v.Name, caseName, v.ComparedName)
}

// StringsAreNotEqual is a validator object.
// Validate adds an error if the Field is equal to ComparedField.
// CaseInsensitive flag can be set to make comparison case insensitive.
type StringsAreNotEqual struct {
	Name            string
	Field           string
	ComparedName    string
	ComparedField   string
	CaseInsensitive bool
	Message         string
}

// Validate adds an error if the Field is equal to ComparedField.
// CaseInsensitive flag can be set to make comparison case insensitive.
func (v *StringsAreNotEqual) Validate(e *validator.Errors) {
	if v.CaseInsensitive {
		if !strings.EqualFold(v.Field, v.ComparedField) {
			return
		}
	} else {
		if v.Field != v.ComparedField {
			return
		}
	}

	e.Add(v.Name, StringsAreNotEqualError(v))
}

// SetField sets validator field.
func (v *StringsAreNotEqual) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringsAreNotEqual) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
