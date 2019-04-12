package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsIPIfaceLocalMulticast(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"ff01::0", true},
		{"fff1::0", true},

		{"127.0.0.1", false},
		{"224.0.0.1", false},

		{"::1", false},
		{"0.0.0.0", false},
		{"feb0::0", false},
		{"169.254.0.0", false},

		{"http://www.google.com", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsIPIfaceLocalMulticast{Name: "IPIface", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsIPIfaceLocalMulticastError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
