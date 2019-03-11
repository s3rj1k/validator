package validators

import (
	"fmt"
	"strings"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsCIDRv4Dive(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"5.255.253.0/24", "220.181.0.0/16", "220.181.0.0/33", "5.255.253.0"}, false, []int{2, 3}},
		{[]string{"fd3b:d101:e37f:9716::/64", "2001:4860:4860::8888/32", "2001:4860:4860::8888/99", "2001:4860:4860::8888"}, false, []int{0, 1, 2, 3}},
		{[]string{" ", ""}, false, []int{0, 1}},
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsCIDRv4{Name: "CIDRv4_Dive"},
			Field:     test.field,
		}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d", index)
		if !test.valid {

			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("CIDRv4_Dive[%d]", i))
			}

			r.Equalf(len(errnames), strings.Count(e.JSON(), "CIDRv4_Dive")/2, "tc %d wrong number of errors", index)

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}
}
