package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// SliceLengthInRangeError is a function that defines error message returned by SliceLengthInRange validator.
// nolint: gochecknoglobals
var SliceLengthInRangeError = func(v *SliceLengthInRange) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	min := v.Min
	max := v.Max

	if max == 0 {
		max = v.length
	}

	if max == -1 {
		return fmt.Sprintf("%v is not empty", v.Field)
	}

	return fmt.Sprintf("%v length=%d not in range(%d, %d)", v.Field, v.length, min, max)
}

// SliceLengthInRange is a validator object.
// Validate adds an error if the slice in Field is not in range between Min and Max (inclusive).
// User can provide either both or only Min/Max fields.
// If only Min provided - Max=length of slice. If only Max provided - Min=0.
// Max=-1 -> slice must be empty.
// Standard check adds an error if the slice is nil or of unsupported type.
type SliceLengthInRange struct {
	Name    string
	Field   interface{}
	Min     int
	Max     int
	Message string
	length  int
}

// Validate adds an error if the slice in Field is not in range between Min and Max (inclusive).
// User can provide either both or only Min/Max fields.
// If only Min provided - Max=length of slice. If only Max provided - Min=0.
// Max=-1 -> slice must be empty.
// Standard check adds an error if the slice is nil or of unsupported type.
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
		return length == 0
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
