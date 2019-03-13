package validators

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsGroupNotExistingDive(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"root", "root ", "wow no"}, false, []int{0}},
		{[]string{" ", ""}, true, []int{}},
		{nil, true, []int{}},
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsGroupNotExisting{Name: "StringIsNotGroup"},
			Field:     test.field,
		}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d", index)
		if !test.valid {
			r.Equalf(len(test.invalidIndexes), e.Count(), "tc %d wrong number of errors", index)

			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("StringIsNotGroup[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}

}
