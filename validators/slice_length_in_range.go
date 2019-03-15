package validators

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// SliceLengthInRangeError is a function that defines error message returned by SliceLengthInRange validator.
// nolint: gochecknoglobals
var SliceLengthInRangeError = func(v *SliceLengthInRange) string {
	if v.Max == -1 {
		return fmt.Sprintf("slice len=%d is not empty", v.length)
	}

	if v.Min == 0 {
		return fmt.Sprintf("slice len=%d is longer than %d", v.length, v.Max)
	}

	if v.Max == 0 {
		return fmt.Sprintf("slice len=%d is shorter than %d", v.length, v.Min)
	}

	return fmt.Sprintf("slice len=%d is not between %d and %d", v.length, v.Min, v.Max)

}

// SliceLengthInRange is a validator object.
type SliceLengthInRange struct {
	Name   string
	Field  interface{}
	Min    int
	Max    int
	length int
}

// Validate adds an error if the slice in Field is not in range between Min and Max (inclusive).
// User can provide either both or only Min/Max fields.
// If only Min provided - Max=length of slice. If only Max provided - Min=0.
// Max=-1 -> slice must be empty
// nolint: gocyclo
func (v *SliceLengthInRange) Validate(e *validator.Errors) {

	if v.Field == nil {
		e.Add(v.Name, ErrNilValue.Error())

		return
	}

	switch field := v.Field.(type) {
	case []string:
		v.length = len(field)
		if lengthOK(len(field), v.Min, v.Max) {
			return
		}
	case []int8:
		v.length = len(field)
		if lengthOK(len(field), v.Min, v.Max) {
			return
		}
	case []int16:
		v.length = len(field)
		if lengthOK(len(field), v.Min, v.Max) {
			return
		}
	case []int32:
		v.length = len(field)
		if lengthOK(len(field), v.Min, v.Max) {
			return
		}
	case []int:
		v.length = len(field)
		if lengthOK(len(field), v.Min, v.Max) {
			return
		}
	case []int64:
		v.length = len(field)
		if lengthOK(len(field), v.Min, v.Max) {
			return
		}
	case []uintptr:
		v.length = len(field)
		if lengthOK(len(field), v.Min, v.Max) {
			return
		}
	case []uint8:
		v.length = len(field)
		if lengthOK(len(field), v.Min, v.Max) {
			return
		}
	case []uint16:
		v.length = len(field)
		if lengthOK(len(field), v.Min, v.Max) {
			return
		}
	case []uint32:
		v.length = len(field)
		if lengthOK(len(field), v.Min, v.Max) {
			return
		}
	case []uint:
		v.length = len(field)
		if lengthOK(len(field), v.Min, v.Max) {
			return
		}
	case []uint64:
		v.length = len(field)
		if lengthOK(len(field), v.Min, v.Max) {
			return
		}
	case []float32:
		v.length = len(field)
		if lengthOK(len(field), v.Min, v.Max) {
			return
		}
	case []float64:
		v.length = len(field)
		if lengthOK(len(field), v.Min, v.Max) {
			return
		}
	case []complex64:
		v.length = len(field)
		if lengthOK(len(field), v.Min, v.Max) {
			return
		}
	case []complex128:
		v.length = len(field)
		if lengthOK(len(field), v.Min, v.Max) {
			return
		}
	default:
		e.Add(v.Name, ErrBadSliceType.Error())
		return
	}

	e.Add(v.Name, SliceLengthInRangeError(v))
}

func lengthOK(length, min, max int) bool {

	if max == -1 {

		if length > 0 {
			return false
		}

		return true
	}

	if min == 0 {
		min = 0
	}

	if max == 0 {
		max = length
	}

	if length >= min && length <= max {
		return true
	}

	return false
}
