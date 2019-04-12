package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// NumbersAreEqualError is a function that defines error message returned by NumbersAreEqual validator.
// nolint: gochecknoglobals
var NumbersAreEqualError = func(v *NumbersAreEqual) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' is not equal to '%s'", NumFieldToString(v.Field), NumFieldToString(v.ComparedField))
	}

	return fmt.Sprintf("'%s' is not equal to '%s'", v.Name, v.ComparedName)
}

// NumbersAreEqual is a validator object.
// Validate adds an error if the Field is not equal to the ComparedField.
type NumbersAreEqual struct {
	Name          string
	Field         interface{}
	ComparedName  string
	ComparedField interface{}
	Message       string
}

// Validate adds an error if the Field is not equal to the ComparedField.
func (v *NumbersAreEqual) Validate(e *validator.Errors) {

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

	if fNum.Value == cfNum.Value {
		return
	}

	e.Add(v.Name, NumbersAreEqualError(v))
}

// SetField sets validator field.
func (v *NumbersAreEqual) SetField(s interface{}) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *NumbersAreEqual) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}

// GetName is a getter on Name field.
func (v *NumbersAreEqual) GetName() string {
	return v.Name
}
