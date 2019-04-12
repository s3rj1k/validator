package buildin

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringHasNoWhitespaceDive(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"3085u", "    a", "z135^^", "()#U	a	]]", "124", "#$%% ", "冊;復製品"}, false, []int{1, 5}}, // 3 has tabs
		{[]string{" ", ""}, false, []int{0}},
		{nil, true, []int{}},
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringHasNoWhitespace{Name: "NoWhitespace"},
			Field:     test.field,
		}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {

			r.Equalf(len(test.invalidIndexes), e.Count(), "tc %d wrong number of errors", index)

			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("NoWhitespace[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}
}
