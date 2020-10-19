package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsIPHasPTR(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"8.8.8.8", true},
		{"208.67.222.222", true},
		{"128.105.39.11", true},
		{"91.198.174.192", true},
		{"72.247.244.88", true},

		{"169.254.0.1", false},
		{"220.181.0.0", false},
		{" 5.255.253.0", false},
		{"http://www.google.com", false},
		{"220.181.0.0/33", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsIPHasPTR{Name: "IP", Field: test.field}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d", index)

		if !test.valid {
			r.Equalf([]string{StringIsIPHasPTRError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
