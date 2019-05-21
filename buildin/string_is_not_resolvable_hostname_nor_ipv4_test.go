package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsNotResolvableHostnameNorIPv4(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"www.google.com", false},
		{"www.i.ua", false},

		{"192.168.0.1", false},
		{"255.255.255.0", false},
		{"8.8.8.8", false},
		{"fd3b:d101:e37f:9716::", true},
		{"0000:0000:0000:0000:0000:0000:0000:0000", true},
		{"2001:db8:a1d5::", true},
		{"random", true},
		{" ", true},
		{"", true},
	}

	for index, test := range tests {
		v := &StringIsNotResolvableHostnameNorIPv4{Name: "RHN", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())
		if !test.valid {
			r.Equalf([]string{StringIsNotResolvableHostnameNorIPv4Error(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
