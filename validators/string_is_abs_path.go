package validators

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsAbsPathError is a function that defines error message returned by StringIsAbsPath validator.
// nolint: gochecknoglobals
var StringIsAbsPathError = func(v *StringIsAbsPath) string {
	return fmt.Sprintf("path '%s' must be absolute", v.Field)
}

// StringIsAbsPath is a validator object
type StringIsAbsPath struct {
	Name  string
	Field string
}

// Validate adds an error if Field is not an absolute path.
func (v *StringIsAbsPath) Validate(e *validator.Errors) {

	if v.Field == "/" {
		return
	}

	if !strings.Contains(v.Field, "//") &&
		strings.TrimSuffix(v.Field, "/") == filepath.Clean(v.Field) &&
		filepath.IsAbs(v.Field) {

		return
	}

	e.Add(v.Name, StringIsAbsPathError(v))
}

// SetField sets validator field.
func (v *StringIsAbsPath) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsAbsPath) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
