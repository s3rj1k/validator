package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_NumbersAreNotEqualDive(t *testing.T) {
	r := require.New(t)

	f := []int{999, 999, 888, 999}
	c := int64(888)

	v := NumberSliceDive{
		Validator: &NumbersAreNotEqual{
			Name:          "MySlice",
			ComparedField: c,
		},
		Field: f,
	}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(1, e.Count())

	c2 := int16(888)

	v = NumberSliceDive{
		Validator: &NumbersAreNotEqual{
			Name:          "MySlice",
			ComparedField: c2,
		},
		Field: f,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())

	f = []int{} // empty array is OK, check in other validator

	v = NumberSliceDive{
		Validator: &NumbersAreNotEqual{
			Name:          "MySlice",
			ComparedField: c,
		},
		Field: f,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	c3 := uint16(888) // cannot pass unsigned and signed

	v = NumberSliceDive{
		Validator: &NumbersAreNotEqual{
			Name:          "MySlice",
			ComparedField: c3,
		},
		Field: f,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	var f2 []int // such array is also ok

	v = NumberSliceDive{
		Validator: &NumbersAreNotEqual{
			Name:          "MySlice",
			ComparedField: c,
		},
		Field: f2,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	f3 := []string{} // fail

	v = NumberSliceDive{
		Validator: &NumbersAreNotEqual{
			Name:          "MySlice",
			ComparedField: c,
		},
		Field: f3,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{ErrBadNumType.Error()}, e.Get("MySlice"))
}
