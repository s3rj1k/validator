package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsIPv6(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"5.255.253.0", false},
		{"220.181.0.0", false},
		{"0.0.0.0", false},
		{"255.255.255.255", false},
		{"fd3b:d101:e37f:9716::", true},
		{"0000:0000:0000:0000:0000:0000:0000:0000", true},
		{"2001:db8:a1d5::", true},

		{"2001:db8:a1d5:: ", false},
		{"http://www.google.com", false},
		{"2001:db8:a1d5::/33", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsIPv6{Name: "IP", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())
		
		if !test.valid {
			r.Equalf([]string{StringIsIPv6Error(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
