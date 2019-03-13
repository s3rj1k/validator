package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsUserExisting(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"root", true},

		{"definitely not", false},
		{"root ", false},

		{" ", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsUserExisting{Name: "GroupE", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d", index)
		if !test.valid {
			r.Equalf([]string{StringIsUserExistingError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
