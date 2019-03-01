package validators

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// NumbersAreNotEqualError is a function that defines error message returned by NumbersAreNotEqual validator.
var NumbersAreNotEqualError = func(v *NumbersAreNotEqual) string {
	return fmt.Sprintf("%d is equal to %d", v.Field, v.ComparedField)
}

// NumbersAreNotEqual is a validator object.
type NumbersAreNotEqual struct {
	Name             string
	Field            interface{}
	ComparedName     string
	ComparedField    interface{}
	ValidateSameType bool
}

// Validate adds an error if the Field is equal to the ComparedField.
func (v *NumbersAreNotEqual) Validate(e *validator.Errors) {

	err := checkNums(v.Field, v.ComparedField, v.ValidateSameType)
	if err != nil {
		e.Add(v.Name, fmt.Sprintf("%s %s", v.Name, err))
		return
	}

	switch field, comparedField := castBoth(v.Field, v.ComparedField); field.(type) {
	case int64:
		if field.(int64) != comparedField.(int64) {
			return
		}
	case uint64:
		if field.(uint64) != comparedField.(uint64) {
			return
		}
	default:
		e.Add(v.Name, fmt.Sprintf("%s %s", v.Name, ErrUnsupportedNumberType))
		return
	}

	e.Add(v.Name, NumbersAreNotEqualError(v))
}

// SetField sets validator field.
func (v *NumbersAreNotEqual) SetField(s interface{}) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *NumbersAreNotEqual) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
