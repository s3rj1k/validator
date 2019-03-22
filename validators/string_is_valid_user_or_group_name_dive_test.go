package validators

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsValidUserOrGroupNameDive(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"_asdgd3y7", "_asdgd3y7$", "asc-af_5522", "x"}, true, []int{}},
		{[]string{"", "1gsdg", "-asfag", "asdasdA", "ajsfn7897$3"}, false, []int{0, 1, 2, 3, 4}},
		{[]string{" ", ""}, false, []int{0, 1}}, // empty is valid
		{nil, false, []int{0}},                  // nil field == empty
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsValidUserOrGroupName{Name: "UGname"},
			Field:     test.field,
		}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf(len(test.invalidIndexes), e.Count(), "tc %d wrong number of errors", index)

			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("UGname[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}
}
