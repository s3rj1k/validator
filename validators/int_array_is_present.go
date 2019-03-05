package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// IntArrayIsPresentError is a function that defines error message returned by IntArrayIsPresent validator.
// nolint: gochecknoglobals
var IntArrayIsPresentError = func(v *IntArrayIsPresent) string {
	return fmt.Sprintf("%s can not be empty", v.Name)
}

// IntArrayIsPresent is a validator object.
type IntArrayIsPresent struct {
	Name  string
	Field []int
}

// Validate adds an error if the field is an empty array.
func (v *IntArrayIsPresent) Validate(e *validator.Errors) {
	if len(v.Field) > 0 {
		return
	}

	e.Add(v.Name, IntArrayIsPresentError(v))
}
