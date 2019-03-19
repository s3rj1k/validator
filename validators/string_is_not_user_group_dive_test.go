package validators

import (
	"fmt"
	"os/user"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsNotUserGroupDive(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"", "baby", "lol"}, true, []int{}},
		{[]string{" ", "", "root"}, true, []int{}}, // empty is invalid
		{nil, true, []int{}},                       // nil field == empty
	}

	// trying to add current user in test if not root
	cu, err := user.Current()
	if err == nil && cu.Gid != "0" {
		tests[0].field = append(tests[0].field, cu.Username)
		tests[0].invalidIndexes = append(tests[0].invalidIndexes, len(tests[0].field)-1)
		tests[0].valid = false
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsNotUserGroup{Name: "UGname"},
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
