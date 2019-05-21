package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsNotExistingUser(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"root", false},

		{"definitely not", true},
		{"root ", true},

		{" ", true},
		{"", true},
	}

	for index, test := range tests {
		v := &StringIsNotExistingUser{Name: "UserNE", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsNotExistingUserError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
