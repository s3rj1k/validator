package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringContainsNoneOf(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field         string
		comparedField []string
		valid         bool
	}{
		{"test", []string{"a", "b", "c"}, true},
		{" test", []string{"well", "done"}, true},
		{"test ", []string{"  ", "mock", "rest"}, true},

		{"test", []string{"e"}, false},
		{" test", []string{" "}, false},
		{"test ", []string{"a", "b", "test"}, false},

		// no required substings
		{"test", []string{}, true},
		{" ", []string{}, true},
		{"test", nil, true},
		{"", nil, true},
	}

	for index, test := range tests {
		v := &StringContainsNoneOf{Name: "StringContains", Field: test.field, ComparedField: test.comparedField}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringContainsNoneOfError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
