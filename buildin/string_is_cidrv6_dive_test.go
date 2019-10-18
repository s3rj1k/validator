package buildin

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsCIDRv6Dive(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"5.255.253.0/24", "220.181.0.0/16", "220.181.0.0/33", "5.255.253.0"}, false, []int{0, 1, 2, 3}},
		{[]string{"fd3b:d101:e37f:9716::/64", "2001:4860:4860::8888/32", "2001:4860:4860::8888/99", "2001:4860:4860::8888"}, false, []int{3}},
		{[]string{" ", ""}, false, []int{0, 1}},
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsCIDRv6{Name: "CIDRv6_Dive"},
			Field:     test.field,
		}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf(len(test.invalidIndexes), e.Count(), "tc %d wrong number of errors", index)

			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("CIDRv6_Dive[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}
}
