package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// NumberIsGreaterError is a function that defines error message returned by NumberIsGreater validator.
// nolint: gochecknoglobals
var NumberIsGreaterError = func(v *NumberIsGreater) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	errt := "is not greater than"

	if v.CheckEqual {
		errt += " or equal to"
	}

	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' %s '%s'", NumFieldToString(v.Field), errt, NumFieldToString(v.ComparedField))
	}

	return fmt.Sprintf("'%s' %s '%s'", v.Name, errt, v.ComparedName)
}

// NumberIsGreater is a validator object.
// Validate adds an error if the Field is not greater than the ComparedField.
type NumberIsGreater struct {
	Name          string
	Field         interface{}
	ComparedName  string
	ComparedField interface{}
	CheckEqual    bool
	Message       string
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

	if v.CheckEqual {
		if fNum.IsGreaterOrEqual(cfNum) {
			return
		}
	} else {
		if fNum.IsGreater(cfNum) {
			return
		}
	}

	e.Add(v.Name, NumberIsGreaterError(v))
}

// SetField sets validator field.
func (v *NumberIsGreater) SetField(s interface{}) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *NumberIsGreater) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}

// GetName is a getter on Name field.
func (v *NumberIsGreater) GetName() string {
	return v.Name
}
