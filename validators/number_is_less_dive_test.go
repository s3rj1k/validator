package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_NumberIsLessDive(t *testing.T) {

	r := require.New(t)

	field := []interface{}{int(10), int8(33), int64(345), int16(-50)}
	compared := 999

	e := validator.NewErrors()
	v := NumberSliceDive{
		Validator: &NumberIsLess{
			Name:          "MySlice",
			ComparedField: compared,
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	e = validator.NewErrors()
	v = NumberSliceDive{
		Validator: &NumberIsLess{
			Name:             "MySlice",
			ComparedField:    compared,
			ValidateSameType: true, // now need only the same type
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(3, e.Count())

	compared = 25

	e = validator.NewErrors()
	v = NumberSliceDive{
		Validator: &NumberIsLess{
			Name:          "MySlice",
			ComparedField: compared,
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(2, e.Count())

	field = []interface{}{nil}

	e = validator.NewErrors()
	v = NumberSliceDive{
		Validator: &NumberIsLess{
			Name:          "MySlice",
			ComparedField: compared,
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(1, e.Count())

	field = []interface{}{"bad type"}
	e = validator.NewErrors()
	v = NumberSliceDive{
		Validator: &NumberIsLess{
			Name:          "MySlice",
			ComparedField: compared,
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(1, e.Count())
}
