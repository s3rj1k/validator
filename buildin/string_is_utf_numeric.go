package buildin

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/s3rj1k/validator"
)

// StringIsUTFNumericError is a function that defines error message returned by StringIsUTFNumeric validator.
// nolint: gochecknoglobals
var StringIsUTFNumericError = func(v *StringIsUTFNumeric) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must contain only unicode numbers", v.Field)
}

// StringIsUTFNumeric is a validator object.
// Validate adds an error if the Field contains anything except unicode numbers (category N).
// Leading sign is allowed. Empty string is valid.
type StringIsUTFNumeric struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field contains anything except unicode numbers (category N).
// Leading sign is allowed. Empty string is valid.
func (v *StringIsUTFNumeric) Validate(e *validator.Errors) {
	var failed bool
	var field = v.Field

	// null string is valid
	if isNullString(field) {
		return
	}

	if strings.IndexAny(field, "+-") > 0 {
		failed = true
	}

	if len(field) > 1 {
		field = strings.TrimPrefix(field, "-")
		field = strings.TrimPrefix(field, "+")
	}

	for _, c := range field {
		if !unicode.IsNumber(c) { //numbers && minus sign are ok
			failed = true
		}
	}

	if failed {
		e.Add(v.Name, StringIsUTFNumericError(v))
	}
}

// SetField sets validator field.
func (v *StringIsUTFNumeric) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsUTFNumeric) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
