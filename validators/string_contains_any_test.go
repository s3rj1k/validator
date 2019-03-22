package validators

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringContainsAny(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field         string
		comparedField []string
		valid         bool
	}{
		{"test", []string{"t", "a", "b"}, true},
		{" test", []string{" t", "well", "done"}, true},
		{"test ", []string{"mock", "mock", "test"}, true},

		{"test", []string{"e ", " s", " ", " t"}, false},
		{" test", []string{"  ", "t "}, false},
		{"test ", []string{"a", "b"}, false},

		// no required substings
		{"test", []string{}, true},
		{" ", []string{}, true},
		{"test", nil, true},
		{"", nil, true},
	}

	for index, test := range tests {
		v := &StringContainsAny{Name: "StringContains", Field: test.field, ComparedField: test.comparedField}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringContainsAnyError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
