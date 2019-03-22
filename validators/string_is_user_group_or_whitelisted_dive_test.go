package validators

import (
	"fmt"
	"os/user"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsUserGroupOrWhitelistedDive(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field          []string
		whitelist      []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"", "baby", "lol"}, []string{"baby", ""}, false, []int{2}},
		{[]string{" ", "", "root"}, nil, false, []int{0, 1, 2}}, // empty is invalid
		{nil, nil, false, []int{0}},                             // nil field == empty
	}

	cu, err := user.Current()
	if err == nil && cu.Gid != "0" {
		tests[0].field[0] = cu.Name
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsUserGroupOrWhitelisted{Name: "UGname", Whitelist: test.whitelist},
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
