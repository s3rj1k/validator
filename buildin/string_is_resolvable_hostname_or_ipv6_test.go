package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsResolvableHostnameOrIPv6(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"www.google.com", true},
		{"www.i.ua", true},
		{"192.168.0.1", false},
		{"255.255.255.0", false},
		{"8.8.8.8", false},
		{"fd3b:d101:e37f:9716::", true},
		{"0000:0000:0000:0000:0000:0000:0000:0000", true},
		{"2001:db8:a1d5::", true},

		{"random", false},
		{" ", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsResolvableHostnameOrIPv6{Name: "RHN", Field: test.field}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsResolvableHostnameOrIPv6Error(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
