package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsHexcolorError is a function that defines error message returned by StringIsHexcolor validator.
// nolint: gochecknoglobals
var StringIsHexcolorError = func(v *StringIsHexcolor) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be a hexadecimal color", v.Field)
}

// StringIsHexcolor is a validator object.
// Validate adds an error if the Field is not formatted as a hexadecimal color.
// Leading '#' is required (e.g. "#1f1f1F", "#F00").
type StringIsHexcolor struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not formatted as a hexadecimal color.
// Leading '#' is required (e.g. "#1f1f1F", "#F00").
func (v *StringIsHexcolor) Validate(e *validator.Errors) {

	if rxHexcolor.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringIsHexcolorError(v))
}

// SetField sets validator field.
func (v *StringIsHexcolor) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsHexcolor) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
