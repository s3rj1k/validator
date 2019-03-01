package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_NumbersAreEqualDive(t *testing.T) {

	r := require.New(t)

	field := []interface{}{int(99), int8(99), int16(99), int32(99)}
	compared := 99

	e := validator.NewErrors()
	v := NumberSliceDive{
		Validator: &NumbersAreEqual{
			Name:          "MySlice",
			ComparedField: compared,
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(0, e.Count())

	e = validator.NewErrors()
	v = NumberSliceDive{
		Validator: &NumbersAreEqual{
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
		Validator: &NumbersAreEqual{
			Name:          "MySlice",
			ComparedField: compared,
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(4, e.Count())

	field = []interface{}{int(25), int8(25), uint16(25), uint32(25)} // cant compare uint and int

	e = validator.NewErrors()
	v = NumberSliceDive{
		Validator: &NumbersAreEqual{
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
		Validator: &NumbersAreEqual{
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
		Validator: &NumbersAreEqual{
			Name:          "MySlice",
			ComparedField: compared,
		},
		Field: field,
	}
	v.Validate(e)
	r.Equal(1, e.Count())
}
