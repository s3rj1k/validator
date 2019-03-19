package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_SliceLengthInRange(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field interface{}
		valid bool
		min   int
		max   int
	}{
		{[]int{10, 11, 55, -10}, true, 1, 4},
		{[]int{10, 11, 55, -10}, true, 0, 0},
		{[]int{10, 11, 55, -10}, false, 0, 1},
		{[]int{10, 11, 55, -10}, false, 5, 8}, // min < max is not checked
		{[]string{"a", "b", "c", "d", ""}, true, 1, 10},
		{[]string{"a", "b", "c", "d", ""}, true, -3, 10}, // min/max > 0 not checked
		{[]string{"a", "b", "c", "d", ""}, false, 0, -1}, // max=-1 -> slice must be empty
		{[]string{}, true, 6, -1},

		{[]rune(""), false, 1, 0},
		{[]rune(" "), true, 1, 0},
		{[]int{}, false, 1, 0},
	}

	for index, test := range tests {
		v := &SliceLengthInRange{Name: "Slice", Field: test.field, Min: test.min, Max: test.max}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())
		
		if !test.valid {
			r.Equal([]string{SliceLengthInRangeError(v)}, e.Get(v.Name))
		}
	}

	v := &SliceLengthInRange{Name: "Slice", Field: nil}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{ErrNilValue.Error()}, e.Get(v.Name))

	v = &SliceLengthInRange{Name: "Slice", Field: []struct{}{}}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal([]string{ErrBadSliceType.Error()}, e.Get(v.Name))
}
