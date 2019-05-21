package buildin

import (
	"fmt"
	"path/filepath"

	"github.com/s3rj1k/validator"
)

// StringsArePathsInTheSameDirError is a function that defines error message returned by StringsArePathsInTheSameDir validator.
// nolint: gochecknoglobals
var StringsArePathsInTheSameDirError = func(v *StringsArePathsInTheSameDir) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' path is not in the same dir with '%s'", v.Field, v.ComparedField)
	}

	return fmt.Sprintf("'%s' path is not in the same dir with '%s'", v.Name, v.ComparedName)
}

// StringsArePathsInTheSameDir is a validator object.
// Validate adds an error if paths do not share same path tree to last path element.
// Supplied paths are converted to absolute paths before comparison.
type StringsArePathsInTheSameDir struct {
	Name          string
	Field         string
	ComparedName  string
	ComparedField string
	Message       string
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

	e.Add(v.Name, StringsArePathsInTheSameDirError(v))
}

// SetField sets validator field.
func (v *StringsArePathsInTheSameDir) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringsArePathsInTheSameDir) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
