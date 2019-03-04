package validators

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringHasNoFileInPathParents is a validator object
type StringHasNoFileInPathParents struct {
	Name  string
	Field string
}

// stringPathToPathParents converts single path to array of absolute path parents
func stringPathToPathParents(path string) []string {

	path, err := filepath.Abs(path)
	if err != nil {
		return nil
	}

	paths := []string{path}

	for len(path) > 1 {
		path = filepath.Dir(path)
		paths = append(paths, path)
	}

	return paths
}

// Validate adds an error if the Field contains path to a file.
func (v *StringHasNoFileInPathParents) Validate(e *validator.Errors) {

	var hasFileInPath bool

	for _, path := range stringPathToPathParents(v.Field) {

		fi, err := os.Stat(path)
		if err != nil {
			continue
		}

		if mode := fi.Mode(); mode.IsDir() {
			continue
		}

		hasFileInPath = true
		break
	}

	if hasFileInPath {
		e.Add(v.Name, fmt.Sprintf("path '%s' contains path to a file", v.Field))
	}
}

// SetField sets validator field.
func (v *StringHasNoFileInPathParents) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringHasNoFileInPathParents) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
