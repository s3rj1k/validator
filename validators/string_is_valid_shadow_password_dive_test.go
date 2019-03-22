package validators

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsValidShadowPasswordDive(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"NP", "x", "!!", "Ep6mckrOLChF.", "$1$X7icamQl$!jSeIrOzj3TQqwTn5kwF/50"}, true, []int{}},
		{[]string{"!!Ep6mckrOLChF.", "X", " !!", " Ep6mckrOLChF.", "$1$X7icamQl$!!jSeIrOzj3TQqwTn5kwF/50"}, false, []int{0, 1, 2, 3, 4}},
		{[]string{" ", ""}, false, []int{0}}, // empty is valid
		{nil, true, []int{}},                 // nil field == empty
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsValidShadowPassword{Name: "ShadowP"},
			Field:     test.field,
		}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf(len(test.invalidIndexes), e.Count(), "tc %d wrong number of errors", index)

			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("ShadowP[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}
}
