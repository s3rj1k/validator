package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsNotNumeric(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"32423a09875", true},
		{"      h", true},
		{"p2#($**)) ", true},
		{"32423A09875", true},
		{"      H", true},
		{"P2#($**)) ", true},

		{"138471938", false},
		{"4", false},
		{"0", false},
		{"0 ", true},
		{"", true},
	}

	for index, test := range tests {
		v := &StringIsNotNumeric{Name: "IsNotNumeric", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%s got=%s", index, !test.valid, e.HasAny())
		if !test.valid {
			r.Equalf([]string{StringIsNotNumericError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
