package buildin

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringsAreEqualError is a function that defines error message returned by StringsAreEqual validator.
// nolint: gochecknoglobals
var StringsAreEqualError = func(v *StringsAreEqual) string {

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
		return fmt.Sprintf("'%s' does not %s '%s'", v.Field, caseName, v.ComparedField)
	}

	return fmt.Sprintf("'%s' does not %s '%s'", v.Name, caseName, v.ComparedName)
}

// StringsAreEqual is a validator object.
// Validate adds an error if the Field is not equal to ComparedField.
// CaseInsensitive flag can be set to make comparison case insensitive.
type StringsAreEqual struct {
	Name            string
	Field           string
	ComparedName    string
	ComparedField   string
	CaseInsensitive bool
	Message         string
}

// Validate adds an error if the Field is not equal to ComparedField.
// CaseInsensitive flag can be set to make comparison case insensitive.
func (v *StringsAreEqual) Validate(e *validator.Errors) {

	if v.CaseInsensitive {
		if strings.EqualFold(v.Field, v.ComparedField) {

			return
		}

	} else {
		if v.Field == v.ComparedField {

			return
		}
	}

	e.Add(v.Name, StringsAreEqualError(v))
}

// SetField sets validator field.
func (v *StringsAreEqual) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringsAreEqual) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
