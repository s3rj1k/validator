package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringHasNoWhitespace(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"32423a09875", true},
		{"      0", false},
		{"2#($**))", true},
		{"32423A09875", true},
		{" 5", false},
		{"P2#($**)) ", false},

		{"abc", true},
		{"$$$%%%%", true},
		{" ", false},
		{"", true},
	}

	for index, test := range tests {
		v := &StringHasNoWhitespace{Name: "NoWhitespace", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringHasNoWhitespaceError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
