package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// NumberIsNotZeroError is a function that defines error message returned by NumberIsNotZero validator.
// nolint: gochecknoglobals
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

	fNum, err := cast(v.Field)
	if err != nil {
		e.Add(v.Name, err.Error())

		return
	}

	if fNum.Value != uint64(0) {
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

// GetName is a getter on Name field.
func (v *NumberIsNotZero) GetName() string {
	return v.Name
}
