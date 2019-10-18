package buildin

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsIPIfaceLocalMulticastDive(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"ff01::0", "ff01:ffff:ffff:ffff:ffff:ffff:ffff:ffff", "fff1::0", "fff1:ffff:ffff:ffff:ffff:ffff:ffff:ffff"}, true, []int{}},
		{[]string{"::1", "0.0.0.0", "feb0::0"}, false, []int{0, 1, 2}},
		{[]string{"127.0.0.1", "224.0.0.2", "8.8.8.8"}, false, []int{0, 1, 2}},
		{[]string{" ", ""}, false, []int{0, 1}},
		{nil, false, []int{0}},
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsIPIfaceLocalMulticast{Name: "IPIface"},
			Field:     test.field,
		}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf(len(test.invalidIndexes), e.Count(), "tc %d wrong number of errors", index)

			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("IPIface[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}
}
