package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_SliceIsNotEmpty(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field interface{}
		valid bool
	}{
		{[]int{10, 11, 55, -10}, true},
		{[]uintptr{1, 11, 1111, 0}, true},
		{[]string{"asd", "qwe", "qwe"}, true},
		{[]float32{0, 15.5}, true},
		{[]string{""}, true},
		{[]byte("here is the byte slice"), true},
		{[]rune{'r', 'u', 'n', 'e'}, true},
		{[]rune("runes"), true},

		{[]rune(""), false},
		{[]byte(""), false},
		{[]int{}, false},
		{[]uintptr{}, false},
		{[]string{}, false},
		{[]byte{}, false},
	}

	for index, test := range tests {
		v := &SliceIsNotEmpty{Name: "Slice", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equal([]string{SliceIsNotEmptyError(v)}, e.Get(v.Name))
		}
	}

	v := &SliceIsNotEmpty{Name: "Slice", Field: nil}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{ErrNilValue.Error()}, e.Get(v.Name))

	v = &SliceIsNotEmpty{Name: "Slice", Field: []struct{}{}}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal([]string{ErrBadSliceType.Error()}, e.Get(v.Name))
}
