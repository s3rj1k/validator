package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsCIDRv6(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"fd3b:d101:e37f:9716::/64", true},
		{"2001:4860:4860::8888/32", true},
		{"2001:4860:4860::8888/99", true},

		{"5.255.253.0/24", false},
		{"220.181.0.0/16", false},
		{"5.255.253.0", false},
		{"2001:4860:4860::8888", false},
		{"220.181.0.0/33", false},
		{" ", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsCIDRv6{Name: "CIDRv4", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d", index)
		if !test.valid {
			r.Equalf([]string{StringIsCIDRv6Error(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
