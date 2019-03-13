package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsNotExistingGroup(t *testing.T) {

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
		v := &StringIsNotExistingGroup{Name: "GroupNE", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d", index)
		if !test.valid {
			r.Equalf([]string{StringIsNotExistingGroupError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
