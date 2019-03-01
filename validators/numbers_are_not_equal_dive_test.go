package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_NumbersAreNotEqualDive(t *testing.T) {

	r := require.New(t)

	field := []interface{}{int(11), int8(11), int16(11), int32(11)}
	compared := 99

	v := NumberSliceDive{
		Validator: &NumbersAreNotEqual{
			Name:          "MySlice",
			ComparedField: compared,
		},
		Field: field,
	}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = NumberSliceDive{
		Validator: &NumbersAreNotEqual{
			Name:             "MySlice",
			ComparedField:    compared,
			ValidateSameType: true, // now need only the same type
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(3, e.Count())

	compared = 11 // now they are equal which is wrong

	v = NumberSliceDive{
		Validator: &NumbersAreNotEqual{
			Name:          "MySlice",
			ComparedField: compared,
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(4, e.Count())

	field = []interface{}{int(25), int8(25), uint16(25), uint32(25)} // not equal but cant compare uint and int

	v = NumberSliceDive{
		Validator: &NumbersAreNotEqual{
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
		Validator: &NumbersAreNotEqual{
			Name:          "MySlice",
			ComparedField: compared,
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())

	field = []interface{}{"bad type"}
	v = NumberSliceDive{
		Validator: &NumbersAreNotEqual{
			Name:          "MySlice",
			ComparedField: compared,
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
}
