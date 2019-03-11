package validators

import (
	"fmt"
	"strings"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsIPv4Dive(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"5.255.253.0", "255.255.255.255", "0.0.0.0", "5.255.253.0/32"}, false, []int{3}},
		{[]string{"fd3b:d101:e37f:9716::", "2001:4860:4860::8888", "2001:4860:4860::8888/99", "2001:4860:4860::8888"}, false, []int{0, 1, 2, 3}},
		{[]string{" ", ""}, false, []int{0, 1}},
		{nil, false, []int{0}},
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsIPv4{Name: "IPv4_Dive"},
			Field:     test.field,
		}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d", index)
		if !test.valid {

			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("IPv4_Dive[%d]", i))
			}

			r.Equalf(len(errnames), strings.Count(e.JSON(), "IPv4_Dive")/2, "tc %d wrong number of errors", index)

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}
}
