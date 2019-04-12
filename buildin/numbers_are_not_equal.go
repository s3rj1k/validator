package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// NumbersAreNotEqualError is a function that defines error message returned by NumbersAreNotEqual validator.
// nolint: gochecknoglobals
var NumbersAreNotEqualError = func(v *NumbersAreNotEqual) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' is equal to '%s'", NumFieldToString(v.Field), NumFieldToString(v.ComparedField))
	}

	return fmt.Sprintf("'%s' is equal to '%s'", v.Name, v.ComparedName)
}

// NumbersAreNotEqual is a validator object.
// Validate adds an error if the Field is equal to the ComparedField.
type NumbersAreNotEqual struct {
	Name          string
	Field         interface{}
	ComparedName  string
	ComparedField interface{}
	Message       string
}

// Validate adds an error if the Field is equal to the ComparedField.
func (v *NumbersAreNotEqual) Validate(e *validator.Errors) {

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

	if fNum.Value != cfNum.Value {
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
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}

// GetName is a getter on Name field.
func (v *NumbersAreNotEqual) GetName() string {
	return v.Name
}
