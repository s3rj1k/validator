package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_NumbersAreEqualDive(t *testing.T) {
	r := require.New(t)

	f := []int{999, 999, 888, 999}
	c := int64(999)

	v := NumberSliceDive{
		Validator: &NumbersAreEqual{
			Name:          "MySlice",
			ComparedField: c,
		},
		Field: f,
	}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(1, e.Count())

	c2 := int16(999)

	v = NumberSliceDive{
		Validator: &NumbersAreEqual{
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
		Validator: &NumbersAreEqual{
			Name:          "MySlice",
			ComparedField: c,
		},
		Field: f,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	c3 := uint16(999) // can pass unsigned and signed

	v = NumberSliceDive{
		Validator: &NumbersAreEqual{
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
		Validator: &NumbersAreEqual{
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
		Validator: &NumbersAreEqual{
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
