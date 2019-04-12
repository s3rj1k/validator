package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// NumberIsNotZeroError is a function that defines error message returned by NumberIsNotZero validator.
// nolint: gochecknoglobals
var NumberIsNotZeroError = func(v *NumberIsNotZero) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must not be equal to 0", v.Name)
}

// NumberIsNotZero is a validator object.
// Validate adds an error if the Field equals to 0.
type NumberIsNotZero struct {
	Name    string
	Field   interface{}
	Message string
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
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}

// GetName is a getter on Name field.
func (v *NumberIsNotZero) GetName() string {
	return v.Name
}
