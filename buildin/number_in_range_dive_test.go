package buildin

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_NumberInRangeDive(t *testing.T) {
	// nolint: maligned
	var tests = []struct {
		checkEqual     bool
		field          interface{}
		min            interface{}
		max            interface{}
		valid          bool
		invalidIndexes []int
	}{
		{
			checkEqual:     false,
			field:          []int{-10, -20, 30},
			min:            -10,
			max:            31,
			valid:          false,
			invalidIndexes: []int{0},
		},
		{true, []uintptr{0, 1, 10, 100}, -100, 100, true, []int{}},

		{false, []uint64{59999, 60000, 60001}, nil, 999999999, true, []int{}},
		{true, []int32{0x1, 1, 200}, 0, 1, false, []int{2}},

		{true, nil, 1, 10, false, []int{0}}, // nil == []int8{0} or 0
		{true, []uint64{1, 2, 3, 4, 5}, nil, nil, false, []int{0, 1, 2, 3, 4}},

		{true, nil, -100, 100, true, []int{}},
		{true, []uint64{1, 2, 3, 4, 5}, nil, 5, true, []int{}},
	}

	r := require.New(t)

	for index, test := range tests {
		v := NumberSliceDive{
			Validator: &NumberInRange{
				Name:       "InRange",
				Min:        test.min,
				Max:        test.max,
				CheckEqual: test.checkEqual},
			Field: test.field,
		}

		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("InRange[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}
}
