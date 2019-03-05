package validators

import (
	"fmt"
	"os"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsNotPath is a validator object
type StringIsNotPath struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is an existing path.
func (v *StringIsNotPath) Validate(e *validator.Errors) {
	if _, err := os.Stat(v.Field); os.IsNotExist(err) {
		return
	}

	e.Add(v.Name, fmt.Sprintf("path '%s' must not exist", v.Field))
}

// SetField sets validator field.
func (v *StringIsNotPath) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNotPath) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
