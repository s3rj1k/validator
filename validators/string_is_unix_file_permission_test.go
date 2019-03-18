package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsUnixFilePermission(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"0750", true},
		{"01750", true},
		{"0777", true},

		{"750", false},
		{"555", false},
		{"999", false},
		{"00777", false},

		{" ", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsUnixFilePermission{Name: "UnixFilePerm", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%s got=%s", index, !test.valid, e.HasAny())
		if !test.valid {
			r.Equalf([]string{StringIsUnixFilePermissionError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
