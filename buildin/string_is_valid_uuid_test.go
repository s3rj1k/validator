package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsValidUUID(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"130dce50-807d-11e9-b475-0800200c9a66", true},
		{"193594C2-807D-11E9-B475-0800200C9A66", true},
		{"24c12192-807d-11e9-bc42-526af7764f64", true},

		{"24c1219-807d-11e9-bc42-526af7764f64", false},
		{"24c12192-807-11e9-bc42-526af7764f64", false},
		{"24c12192-807d-11e-bc42-526af7764f64", false},
		{"24c12192-807d-11e9-bc4-526af7764f64", false},
		{"24c12192-807d-11e9-bc42-526af7764f6", false},

		{" 24c12192-807d-11e9-bc42-526af7764f64", false},
		{"24c12192-807d-11e9- bc42-526af7764f64", false},
		{"24c12192-807d-11e9-bc42-526af7764f64 ", false},

		{"g4c12192-807d-11e9-bc42-526af7764f64", false},

		{"random", false},
		{" ", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsValidUUID{Name: "UUID", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())
		if !test.valid {
			r.Equalf([]string{StringIsValidUUIDError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
