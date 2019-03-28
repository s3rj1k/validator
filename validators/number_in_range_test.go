package validators

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_NumberInRange(t *testing.T) {

	// nolint: maligned
	var tests = []struct {
		checkEqual bool
		field      interface{}
		min        interface{}
		max        interface{}
		valid      bool
	}{
		{
			checkEqual: false,
			field:      -10,
			min:        -10,
			max:        31,
			valid:      false,
		},
		{
			checkEqual: true,
			field:      -10,
			min:        -10,
			max:        31,
			valid:      true,
		},
		{true, uintptr(20), -100, 100, true},

		{false, 1, nil, 999999999, true}, // nil == 0
		{false, int64(1), 0, 1, false},

		{true, nil, 0, 10, true}, // nil == 0
		{true, 1, nil, nil, false},

		{true, nil, -100, 100, true},
		{false, 0, nil, 5, false},
	}

	r := require.New(t)

	for index, tc := range tests {

		v := &NumberInRange{
			Name:       "InRange",
			Field:      tc.field,
			CheckEqual: tc.checkEqual,
			Min:        tc.min,
			Max:        tc.max,
		}

		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!tc.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !tc.valid, e.HasAny())

		if !tc.valid {
			r.Equalf([]string{NumberInRangeError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
