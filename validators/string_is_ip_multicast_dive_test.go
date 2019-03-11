package validators

import (
	"fmt"
	"strings"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsIPMulticastDive(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"ff00::1", "ffff::0", "224.0.0.0", "239.255.255.255"}, true, []int{}},
		{[]string{"2000::0", "3fff::0", "8.8.8.8", "172.16.0.0"}, false, []int{0, 1, 2, 3}},
		{[]string{"127.0.0.0", "0.0.0.0"}, false, []int{0, 1}},
		{nil, false, []int{0}},
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsIPMulticast{Name: "IPMulticast"},
			Field:     test.field,
		}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d", index)
		if !test.valid {

			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("IPMulticast[%d]", i))
			}

			r.Equalf(len(errnames), strings.Count(e.JSON(), "IPMulticast")/2, "tc %d wrong number of errors", index)

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}
}
