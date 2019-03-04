package validators

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_NumbersAreNotEqual(t *testing.T) {

	r := require.New(t)

	f := int32(90)
	c := int32(100)

	v := &NumbersAreNotEqual{Name: "Number", Field: f, ComparedField: c}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	c2 := int16(100) // this also works

	v = &NumbersAreNotEqual{Name: "Number", Field: f, ComparedField: c2}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	c3 := uint(100) // this also works

	v = &NumbersAreNotEqual{Name: "Number", Field: f, ComparedField: c3}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	f2 := []interface{}{} // obviously not

	for _, i := range f2 {
		v := &NumbersAreNotEqual{Name: "Number", Field: i, ComparedField: c3}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(1, e.Count())
		r.Equal([]string{ErrBadNumType.Error()}, e.Get("Number"))
	}

	f3 := []string{"world"} // also not

	for _, i := range f3 {
		v := &NumbersAreNotEqual{Name: "Number", Field: i, ComparedField: c3}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(1, e.Count())
		r.Equal([]string{ErrBadNumType.Error()}, e.Get("Number"))
	}
}
