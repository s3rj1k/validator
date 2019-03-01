package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_NumberIsGreaterDive(t *testing.T) {

	r := require.New(t)

	field := []interface{}{int(10), int8(33), int64(345), int16(-50)}
	compared := -100

	v := NumberSliceDive{
		Validator: &NumberIsGreater{
			Name:          "MySlice",
			ComparedField: compared,
		},
		Field: field,
	}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = NumberSliceDive{
		Validator: &NumberIsGreater{
			Name:             "MySlice",
			ComparedField:    compared,
			ValidateSameType: true, // now need only the same type
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(3, e.Count())

	compared = 25

	v = NumberSliceDive{
		Validator: &NumberIsGreater{
			Name:          "MySlice",
			ComparedField: compared,
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(2, e.Count())

	field = []interface{}{nil}

	v = NumberSliceDive{
		Validator: &NumberIsGreater{
			Name:          "MySlice",
			ComparedField: compared,
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())

	field = []interface{}{"bad type"}
	e = validator.NewErrors()
	v = NumberSliceDive{
		Validator: &NumberIsGreater{
			Name:          "MySlice",
			ComparedField: compared,
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(1, e.Count())
}
