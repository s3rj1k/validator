package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_NumbersAreEqual(t *testing.T) {
	r := require.New(t)

	f := []int32{100, 100, 100}
	c := int32(100)

	for _, i := range f {
		v := &NumbersAreEqual{Name: "Number", Field: i, ComparedField: c}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equal(0, e.Count())
	}

	c2 := int16(100) // this also works

	for _, i := range f {
		v := &NumbersAreEqual{Name: "Number", Field: i, ComparedField: c2}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equal(0, e.Count())
	}

	c3 := uint(100) // this will not

	for _, i := range f {
		v := &NumbersAreEqual{Name: "Number", Field: i, ComparedField: c3}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equal(0, e.Count())
	}

	f2 := []interface{}{} // obviously not

	for _, i := range f2 {
		v := &NumbersAreEqual{Name: "Number", Field: i, ComparedField: c3}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equal(1, e.Count())
		r.Equal([]string{ErrBadNumType.Error()}, e.Get("Number"))
	}

	f3 := []string{"world"} // also not

	for _, i := range f3 {
		v := &NumbersAreEqual{Name: "Number", Field: i, ComparedField: c3}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equal(1, e.Count())
		r.Equal([]string{ErrBadNumType.Error()}, e.Get("Number"))
	}
}
