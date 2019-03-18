package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsRGBcolorError is a function that defines error message returned by StringIsRGBcolor validator.
// nolint: gochecknoglobals
var StringIsRGBcolorError = func(v *StringIsRGBcolor) string {
	return fmt.Sprintf("'%s' must be a RGB color in format rgb(RRR, GGG, BBB)", v.Field)
}

// StringIsRGBcolor is a validator object.
type StringIsRGBcolor struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not formatted as an RGB color.
// Expected format is "rgb(RRR, GGG, BBB)".
func (v *StringIsRGBcolor) Validate(e *validator.Errors) {

	if rxRGBcolor.MatchString(v.Field) {
		return
	}

	e.Add(v.Name, StringIsRGBcolorError(v))
}

// SetField sets validator field.
func (v *StringIsRGBcolor) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsRGBcolor) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
