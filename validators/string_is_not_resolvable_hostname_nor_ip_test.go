package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsNotResolvableHostnameNorIP(t *testing.T) {

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
		{"fd3b:d101:e37f:9716::", false},
		{"0000:0000:0000:0000:0000:0000:0000:0000", false},
		{"2001:db8:a1d5::", false},
		{"random", true},
		{" ", true},
		{"", true},
	}

	for index, test := range tests {
		v := &StringIsNotResolvableHostnameNorIP{Name: "RHN", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())
		if !test.valid {
			r.Equalf([]string{StringIsNotResolvableHostnameNorIPError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
