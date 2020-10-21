package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_NumberInSlice(t *testing.T) {
	// nolint: maligned
	var tests = []struct {
		field interface{}
		slice interface{}
		valid bool
	}{
		{
			field: -10,
			slice: []int32{3, -10, 8888},
			valid: true,
		},
		{
			field: -8888,
			slice: []int32{3, -10, 8888},
			valid: false,
		},
		{int64(1), []uint32{1, 2}, true},
		{int64(-1), []uint32{1, 2}, false},

		{0, []int32{0, 1, 2}, true},
		{nil, []int32{1, 2}, false}, // nil == 0

		{nil, []int32{0, 1, 2}, true}, // nil == 0
		{nil, nil, false},             // nil == 0
	}

	r := require.New(t)

	for index, tc := range tests {
		v := &NumberInSlice{
			Name:  "InSlice",
			Field: tc.field,
			Slice: tc.slice,
		}

		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!tc.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !tc.valid, e.HasAny())

		if !tc.valid {
			r.Equalf([]string{NumberInSliceError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
