package validators

import (
	"fmt"
	"path/filepath"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringsArePathsInTheSameDir is a validator object
type StringsArePathsInTheSameDir struct {
	Name          string
	Field         string
	ComparedName  string
	ComparedField string
}

// Validate adds an error if paths do not share same path tree to last path element.
// Supplied paths are converted to absolute paths before comparison.
func (v *StringsArePathsInTheSameDir) Validate(e *validator.Errors) {

	absFieldPath, _ := filepath.Abs(v.Field)
	absComparedFieldPath, _ := filepath.Abs(v.ComparedField)

	if (filepath.Dir(absFieldPath) == filepath.Dir(absComparedFieldPath)) &&
		filepath.IsAbs(absFieldPath) &&
		filepath.IsAbs(absComparedFieldPath) {
		return
	}

	if len(v.ComparedName) == 0 {
		e.Add(v.Name, fmt.Sprintf("'%s' path is not in the same dir with '%s'", v.Name, v.ComparedField))
	} else {
		e.Add(v.Name, fmt.Sprintf("'%s' path is not in the same dir with '%s'", v.Name, v.ComparedName))
	}
}

// SetField sets validator field.
func (v *StringsArePathsInTheSameDir) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringsArePathsInTheSameDir) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
