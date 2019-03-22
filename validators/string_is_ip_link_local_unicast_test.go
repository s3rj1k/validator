package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsIPLinkLocalUnicast(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"fe80::0", true},
		{"febf::ffff", true},

		{"169.254.0.0", true},
		{"169.254.255.255", true},

		{"ff01::0", false},
		{"ff02::0", false},

		{"::1", false},
		{"0.0.0.0", false},
		{"fed0::0", false},

		{"http://www.google.com", false},
		{"8.8.8.8", false},
	}

	for index, test := range tests {
		v := &StringIsIPLinkLocalUnicast{Name: "IPLink", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsIPLinkLocalUnicastError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
