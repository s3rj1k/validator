package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsIPLinkLocalMulticast(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"ff02::0", true},
		{"fff2::0", true},

		{"224.0.0.0", true},
		{"224.0.0.255", true},

		{"::1", false},
		{"0.0.0.0", false},
		{"feb0::0", false},
		{"169.254.0.0", false},

		{"http://www.google.com", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsIPLinkLocalMulticast{Name: "IPLink", Field: test.field}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsIPLinkLocalMulticastError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
