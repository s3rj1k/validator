package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// NumberIsNotZeroError is a function that defines error message returned by NumberIsNotZero validator.
var NumberIsNotZeroError = func(v *NumberIsNotZero) string {
	return fmt.Sprintf("%s must not be equal to 0", v.Name)
}

// NumberIsNotZero is a validator object.
type NumberIsNotZero struct {
	Name  string
	Field interface{}
}

// Validate adds an error if the Field equals to 0.
func (v *NumberIsNotZero) Validate(e *validator.Errors) {

	if isNil(v.Field) {
		e.Add(v.Name, fmt.Sprintf("%s %s", v.Name, ErrNilFields))
		return
	}

	value, _, _ := cast(v.Field)

	switch value.(type) {
	case int64:
		if value.(int64) != int64(0) {
			return
		}
	case uint64:
		if value.(uint64) != uint64(0) {
			return
		}
	default:
		e.Add(v.Name, fmt.Sprintf("%s %s", v.Name, ErrUnsupportedNumberType))
		return
	}

	e.Add(v.Name, NumberIsNotZeroError(v))
}

// SetField sets validator field.
func (v *NumberIsNotZero) SetField(s interface{}) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *NumberIsNotZero) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
