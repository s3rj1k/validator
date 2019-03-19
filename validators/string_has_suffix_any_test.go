package validators

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringHasSuffixAny(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field         string
		comparedField []string
		valid         bool
	}{
		{"string", []string{"abc", "def", "ing"}, true},
		{"string ", []string{"a", "b", "ing "}, true},
		{" ", []string{""}, true},
		{"", []string{}, true},
		{"string", nil, true},
		{"", nil, true},

		{"string", []string{"abc", "def", " ing"}, false},
		{"string ", []string{"abc", "def", "ing"}, false},
	}

	for index, test := range tests {
		v := &StringHasSuffixAny{Name: "StringSuffix", Field: test.field, ComparedField: test.comparedField}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())
		
		if !test.valid {
			r.Equalf([]string{StringHasSuffixAnyError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}

}
