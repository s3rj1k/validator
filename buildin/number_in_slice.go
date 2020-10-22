package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// NumberInSliceError is a function that defines error message returned by NumberInSlice validator.
// nolint: gochecknoglobals
var NumberInSliceError = func(v *NumberInSlice) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	e := fmt.Sprintf("'%s' not in slice %v", NumFieldToString(v.Field), v.Slice)

	return e
}

// NumberInSlice is a validator object.
// Validate adds an error if the Field is not in slice.
type NumberInSlice struct {
	Name    string
	Field   interface{}
	Slice   interface{}
	Message string
}

// Validate adds an error if the Field is not in slice.
// Empty Field value will be treated as 0 (zero).
// Empty Slice value will be treated as empty []*Number slice.
func (v *NumberInSlice) Validate(e *validator.Errors) {
	fNum, err := cast(v.Field)
	if err != nil {
		e.Add(v.Name, err.Error())

		return
	}

	fSlice, err := castSlice(v.Slice)
	if err != nil {
		e.Add(v.Name, err.Error())

		return
	}

	if fNum.InSlice(fSlice) {
		return
	}

	e.Add(v.Name, NumberInSliceError(v))
}

// SetField sets validator field.
func (v *NumberInSlice) SetField(s interface{}) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *NumberInSlice) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}

// GetName is a getter on Name field.
func (v *NumberInSlice) GetName() string {
	return v.Name
}
