package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// NumberInRangeError is a function that defines error message returned by NumberInRange validator.
// nolint: gochecknoglobals
var NumberInRangeError = func(v *NumberInRange) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	e := fmt.Sprintf("'%s' not in range(%s, %s)", NumFieldToString(v.Field), NumFieldToString(v.Min), NumFieldToString(v.Max))

	if v.CheckEqual {
		return fmt.Sprintf("%s (inclusive)", e)
	}

	return e
}

// NumberInRange is a validator object.
// Validate adds an error if the Field is not in range between Min and Max (inclusive).
// Empty Min/Max values will be treated as 0 (zeros).
type NumberInRange struct {
	Name       string
	Field      interface{}
	Min        interface{}
	Max        interface{}
	CheckEqual bool
	Message    string
}

// Validate adds an error if the Field is not in range between Min and Max (inclusive).
// Empty Min/Max values will be treated as 0 (zeros).
func (v *NumberInRange) Validate(e *validator.Errors) {
	fNum, err := cast(v.Field)
	if err != nil {
		e.Add(v.Name, err.Error())

		return
	}

	min, err := cast(v.Min)
	if err != nil {
		e.Add(v.Name, err.Error())

		return
	}

	max, err := cast(v.Max)
	if err != nil {
		e.Add(v.Name, err.Error())

		return
	}

	if v.CheckEqual {
		if fNum.InRangeOrEqual(min, max) {
			return
		}
	} else {
		if fNum.InRange(min, max) {
			return
		}
	}

	e.Add(v.Name, NumberInRangeError(v))
}

// SetField sets validator field.
func (v *NumberInRange) SetField(s interface{}) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *NumberInRange) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}

// GetName is a getter on Name field.
func (v *NumberInRange) GetName() string {
	return v.Name
}
