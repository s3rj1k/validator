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
		isErr            bool
		duplicateIndexes []int
	}{
		{[]int{10, 11, 55, -10}, false, []int{}},
		{[]int{99, 100, 101, 100}, true, []int{3}},
		{[]uintptr{1, 11, 1111, 0}, false, []int{}},
		{[]uintptr{123, 321, 123, 321}, true, []int{2, 3}},
		{[]int{}, false, []int{}},
		{[]uintptr{}, false, []int{}},
		{[]string{"hello", "world"}, false, []int{}},
		{[]string{"hello", "world", "hello", "hello"}, true, []int{2, 3}},
	}

	for _, test := range tests {
		v := &SliceIsUnique{Name: "Slice", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equal(test.isErr, e.HasAny())
		if test.isErr {
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
