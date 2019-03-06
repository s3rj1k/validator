package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// NumberIsLessError is a function that defines error message returned by NumberIsLess validator.
// nolint: gochecknoglobals
var NumberIsLessError = func(v *NumberIsLess) string {
	return fmt.Sprintf("%d is not less than %d", v.Field, v.ComparedField)
}

// NumberIsLess is a validator object.
type NumberIsLess struct {
	Name          string
	Field         interface{}
	ComparedName  string
	ComparedField interface{}
	CheckEqual    bool
}

// Validate adds an error if the Field is not less than the ComparedField.
func (v *NumberIsLess) Validate(e *validator.Errors) {

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

	if isLess(fNum, cfNum, v.CheckEqual) {
		return
	}

	e.Add(v.Name, NumberIsLessError(v))
}

// SetField sets validator field.
func (v *NumberIsLess) SetField(s interface{}) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *NumberIsLess) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}

// GetName is a getter on Name field.
func (v *NumberIsLess) GetName() string {
	return v.Name
}

// isLess returns true if x < y or x<=y if checkEqual is true
func isLess(x, y *Number, checkEqual bool) bool {

	switch {
	case x.isNegative && !y.isNegative:
		return true
	case !x.isNegative && y.isNegative:
		return false

	case !x.isNegative && !y.isNegative && checkEqual:
		return x.Value <= y.Value
	case !x.isNegative && !y.isNegative && !checkEqual:
		return x.Value < y.Value

	case x.isNegative && y.isNegative && checkEqual:
		return x.Value >= y.Value
	case x.isNegative && y.isNegative && !checkEqual:
		return x.Value > y.Value
	}

	return false
}
