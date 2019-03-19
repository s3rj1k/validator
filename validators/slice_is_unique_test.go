package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_SliceIsUnique(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field            interface{}
		valid            bool
		duplicateIndexes []int
	}{
		{[]int{10, 11, 55, -10}, true, []int{}},
		{[]int{99, 100, 101, 100}, false, []int{3}},
		{[]uintptr{1, 11, 1111, 0}, true, []int{}},
		{[]uintptr{123, 321, 123, 321}, false, []int{2, 3}},
		{[]int{}, true, []int{}},
		{[]uintptr{}, true, []int{}},
		{[]string{"hello", "world"}, true, []int{}},
		{[]string{"hello", "world", "hello", "hello"}, false, []int{2, 3}},
	}

	for index, test := range tests {
		v := &SliceIsUnique{Name: "Slice", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equal([]string{SliceIsUniqueError(v)}, e.Get(v.Name))
		}
	}

	v := &SliceIsUnique{Name: "Slice", Field: nil}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{ErrNilValue.Error()}, e.Get(v.Name))

	v = &SliceIsUnique{Name: "Slice", Field: []struct{}{}}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal([]string{ErrBadSliceType.Error()}, e.Get(v.Name))
}
