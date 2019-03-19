package validators

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringHasPrefixAnyDive(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field          []string
		comparedField  []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"foo", "bar", "golang "}, []string{"f", "b", "g"}, true, nil},
		{[]string{" bar  ", "bar", ""}, []string{"", " ", " b"}, true, nil},

		{[]string{}, []string{""}, true, nil},
		{[]string{"foo", "bar  ", " golang"}, []string{""}, true, nil},
		{[]string{}, []string{}, true, nil},
		{nil, []string{}, true, nil},
		{[]string{}, nil, true, nil},
		{nil, nil, true, nil},

		{[]string{"foo", "bar", "golang ", ""}, []string{"d", "a", "b"}, false, []int{0, 2, 3}},
		{[]string{"foo", " bar  ", "bar", ""}, []string{"d", "a", " b"}, false, []int{0, 2, 3}},
		{[]string{"foo", "bar", " golang ", ""}, []string{" "}, false, []int{0, 1, 3}},
	}

	for index, test := range tests {
		v := StringSliceDive{
			Validator: &StringHasPrefixAny{Name: "StringPrefix", ComparedField: test.comparedField},
			Field:     test.field,
		}

		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())
		
		if !test.valid {
			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("StringPrefix[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}

}
