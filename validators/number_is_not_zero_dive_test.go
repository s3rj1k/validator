package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_NumberIsNotZeroDive(t *testing.T) {

	r := require.New(t)

	field := []interface{}{int(10), int8(33), int64(345), int16(-50)}

	e := validator.NewErrors()
	v := NumberSliceDive{
		Validator: &NumberIsNotZero{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []interface{}{int(0), int8(0), uint64(0), uintptr(0)}

	e = validator.NewErrors()
	v = NumberSliceDive{
		Validator: &NumberIsNotZero{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(4, e.Count())

	field = []interface{}{nil}

	e = validator.NewErrors()
	v = NumberSliceDive{
		Validator: &NumberIsNotZero{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(1, e.Count())

	field = []interface{}{"bad type"}
	e = validator.NewErrors()
	v = NumberSliceDive{
		Validator: &NumberIsNotZero{
			Name: "MySlice",
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(1, e.Count())
}
