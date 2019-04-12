package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringHasPrefixNoneOf(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field         string
		comparedField []string
		valid         bool
	}{
		{"test", []string{"e", "s", " ", " t"}, true},
		{" test", []string{" te ", "t ", "t"}, true},
		{"test ", []string{" ", " t"}, true},
		{"test", []string{}, true},
		{" ", []string{}, true},
		{"test", nil, true},
		{"", nil, true},

		{"test", []string{"t"}, false},
		{" test", []string{"mock", "mock", " te"}, false},
		{"test ", []string{"mock", "mock", "test "}, false},
	}

	for index, test := range tests {
		v := &StringHasPrefixNoneOf{Name: "StringPrefix", Field: test.field, ComparedField: test.comparedField}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringHasPrefixNoneOfError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
