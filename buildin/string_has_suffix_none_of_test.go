package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringHasSuffixNoneOf(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field         string
		comparedField []string
		valid         bool
	}{
		{"string", []string{"g ", "in", " "}, true},
		{"string ", []string{"ing"}, true},
		{"", []string{" "}, true},
		{"", []string{}, true},
		{"string", nil, true},
		{"", nil, true},

		{"string", []string{"ing"}, false},
		{"string ", []string{"g "}, false},
		{"string ", []string{"string "}, false},
		{"string", []string{""}, false},
		{"string ", []string{" "}, false},
	}

	for index, test := range tests {
		v := &StringHasSuffixNoneOf{Name: "StringSuffix", Field: test.field, ComparedField: test.comparedField}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringHasSuffixNoneOfError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
