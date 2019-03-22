package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsIPGlobalUnicast(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"2000::", true},
		{"3fff:ffff:ffff:ffff:ffff:ffff:ffff:ffff", true},
		{"2600:3c00::f03c:91ff:fe67:aa7c", true},

		{"8.8.8.8", true},
		{"209.185.108.134", true},

		{"172.16.0.0", true},  // these are allowed
		{"192.168.0.0", true}, // per net package
		{"fec0::0", true},     // ip.IsGlobalUnicast() func

		{"::1", false},
		{"0.0.0.0", false},
		{"feb0::0", false},
		{"169.254.0.0", false},

		{"http://www.google.com", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsIPGlobalUnicast{Name: "IP", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsIPGlobalUnicastError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
