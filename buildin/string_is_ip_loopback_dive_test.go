package buildin

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsIPLoopbackDive(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"::1", "127.0.0.0", "127.255.255.255"}, true, []int{}},
		{[]string{"2000::0", "3fff::0", "8.8.8.8", "172.16.0.0"}, false, []int{0, 1, 2, 3}},
		{[]string{" ", ""}, false, []int{0, 1}},
		{nil, false, []int{0}},
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsIPLoopback{Name: "IPLoopback"},
			Field:     test.field,
		}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf(len(test.invalidIndexes), e.Count(), "tc %d wrong number of errors", index)

			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("IPLoopback[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}
}
