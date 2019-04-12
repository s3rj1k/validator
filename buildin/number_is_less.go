package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// NumberIsLessError is a function that defines error message returned by NumberIsLess validator.
// nolint: gochecknoglobals
var NumberIsLessError = func(v *NumberIsLess) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	errt := "is not less than"

	if v.CheckEqual {
		errt += " or equal to"
	}

	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' %s '%s'", NumFieldToString(v.Field), errt, NumFieldToString(v.ComparedField))
	}

	return fmt.Sprintf("'%s' %s '%s'", v.Name, errt, v.ComparedName)
}

// NumberIsLess is a validator object.
// Validate adds an error if the Field is not less than the ComparedField.
type NumberIsLess struct {
	Name          string
	Field         interface{}
	ComparedName  string
	ComparedField interface{}
	CheckEqual    bool
	Message       string
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
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}

// GetName is a getter on Name field.
func (v *NumberIsLess) GetName() string {
	return v.Name
}

// isLess returns true if x < y or x<=y if checkEqual is true
func isLess(x, y *Number, checkEqual bool) bool {

	if x.IsLess(y) {
		return true
	}

	if checkEqual && x.IsEqual(y) {
		return true
	}

	return false
}
