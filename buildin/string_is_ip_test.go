package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsIP(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"5.255.253.0", true},
		{"220.181.0.0", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"fd3b:d101:e37f:9716::", true},
		{"0000:0000:0000:0000:0000:0000:0000:0000", true},
		{"2001:db8:a1d5::", true},

		{" 5.255.253.0", false},
		{"http://www.google.com", false},
		{"220.181.0.0/33", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsIP{Name: "IP", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d", index)
		if !test.valid {
			r.Equalf([]string{StringIsIPError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
