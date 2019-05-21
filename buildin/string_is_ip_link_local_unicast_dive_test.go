package buildin

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsIPLinkLocalUnicastDive(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"fe80::0", "fe80:ffff:ffff:ffff:ffff:ffff:ffff:ffff", "febf::0", "febf:ffff:ffff:ffff:ffff:ffff:ffff:ffff", "169.254.0.0"}, true, []int{}},
		{[]string{"224.0.0.0", "224.0.0.255"}, false, []int{0, 1}},
		{[]string{"127.0.0.1", "::1", "0.0.0.0", "feb0::0 "}, false, []int{0, 1, 2, 3}},
		{[]string{" ", ""}, false, []int{0, 1}},
		{nil, false, []int{0}},
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsIPLinkLocalUnicast{Name: "IPLink"},
			Field:     test.field,
		}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf(len(test.invalidIndexes), e.Count(), "tc %d wrong number of errors", index)

			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("IPLink[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}
}
