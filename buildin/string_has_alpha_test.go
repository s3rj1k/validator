package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringHasAlpha(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"32423a09875", true},
		{"      h", true},
		{"p2#($**)) ", true},
		{"32423A09875", true},
		{"      H", true},
		{"P2#($**)) ", true},

		{"138471938", false},
		{"$$$%%%%", false},
		{" ", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringHasAlpha{Name: "HasAlpha", Field: test.field}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringHasAlphaError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
