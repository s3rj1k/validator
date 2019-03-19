package validators

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringHasPrefixAny(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field         string
		comparedField []string
		valid         bool
	}{
		{"test", []string{"t", "dont", "matter"}, true},
		{" test", []string{"mock", "mock", " te"}, true},
		{"test ", []string{"mock", "mock", "test "}, true},

		{"test", []string{"e", "s", " ", " t"}, false},
		{" test", []string{" te ", "t ", "t"}, false},
		{"test ", []string{" ", " t"}, false},

		// no required prefixes
		{"test", []string{}, true},
		{" ", []string{}, true},
		{"test", nil, true},
		{"", nil, true},
	}

	for index, test := range tests {
		v := &StringHasPrefixAny{Name: "StringPrefix", Field: test.field, ComparedField: test.comparedField}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())
		
		if !test.valid {
			r.Equalf([]string{StringHasPrefixAnyError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
