package buildin

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_NumberInSliceDive(t *testing.T) {
	// nolint: maligned
	var tests = []struct {
		field          interface{}
		slice          interface{}
		valid          bool
		invalidIndexes []int
	}{
		{
			field:          []int{-10, -8888, 8888},
			slice:          []int32{3, -10, 8888},
			valid:          false,
			invalidIndexes: []int{1},
		},
		{[]int64{-1, 1}, []uint32{1, 2}, false, []int{0}},
		{[]int32{0x1, 0}, []uint32{1, 2}, false, []int{}},

		{nil, []int32{1, 2}, false, []int{}}, // nil == []int8{0} or 0
		{[]uint64{1, 2, 3, 4, 5}, nil, false, []int{}},

		{nil, []uint32{0, 1, 2}, true, []int{}},
		{nil, nil, false, []int{}},
	}

	r := require.New(t)

	for index, test := range tests {
		v := NumberSliceDive{
			Validator: &NumberInSlice{
				Name:  "InSlice",
				Slice: test.slice},
			Field: test.field,
		}

		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("InSlice[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}
}
