package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// NumberIsGreaterError is a function that defines error message returned by NumberIsGreater validator.
// nolint: gochecknoglobals
var NumberIsGreaterError = func(v *NumberIsGreater) string {
	return fmt.Sprintf("%d is not greater than %d", v.Field, v.ComparedField)
}

// NumberIsGreater is a validator object.
type NumberIsGreater struct {
	Name          string
	Field         interface{}
	ComparedName  string
	ComparedField interface{}
	CheckEqual    bool
}

// Validate adds an error if the Field is not greater than the ComparedField.
func (v *NumberIsGreater) Validate(e *validator.Errors) {

	fNum, err := cast(v.Field)
	if err != nil {
		e.Add(v.Name, err.Error())

		return
	}

	cfNum, err := cast(v.ComparedField)
	if err != nil {
		e.Add(v.Name, err.Error())

		return
	}

	if isGreater(fNum, cfNum, v.CheckEqual) {
		return
	}

	e.Add(v.Name, NumberIsGreaterError(v))
}

// SetField sets validator field.
func (v *NumberIsGreater) SetField(s interface{}) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *NumberIsGreater) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}

// GetName is a getter on Name field.
func (v *NumberIsGreater) GetName() string {
	return v.Name
}

// isGreater returns true if x > y or x>=y if checkEqual is true
func isGreater(x, y *Number, checkEqual bool) bool {

	switch {
	case x.isNegative && !y.isNegative:
		return false
	case !x.isNegative && y.isNegative:
		return true

	case !x.isNegative && !y.isNegative && checkEqual:
		return x.Value >= y.Value
	case !x.isNegative && !y.isNegative && !checkEqual:
		return x.Value > y.Value

	case x.isNegative && y.isNegative && checkEqual:
		return x.Value <= y.Value
	case x.isNegative && y.isNegative && !checkEqual:
		return x.Value < y.Value
	}

	return false
}
