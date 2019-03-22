package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsIPLoopback(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"::1", true},
		{"127.0.0.0", true},
		{"127.255.255.255", true},

		{"8.8.8.8", false},
		{"209.185.108.134", false},

		{"0.0.0.0", false},
		{"feb0::0", false},
		{"169.254.0.0", false},

		{"http://www.google.com", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsIPLoopback{Name: "IP", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsIPLoopbackError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
