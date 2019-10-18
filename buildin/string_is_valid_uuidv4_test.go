package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsValidUUIDv4(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"2a72e244-2cee-4ede-a0d6-cdf38098f453", true},
		{"1DDF8143-9EEE-4E1B-8410-3F48DED2FB73", true},
		{"4a84dfca-43e0-46ea-a7bb-dad298ce5ded", true},

		{"4a84dfca-43e0-46ea-b7bb-dad298ce5ded", true},
		{"4a84dfca-43e0-46ea-87bb-dad298ce5ded", true},
		{"4a84dfca-43e0-46ea-97bb-dad298ce5ded", true},

		{"4a84dfca-43e0-36ea-a7bb-dad298ce5ded", false},
		{"4a84dfca-43e0-46ea-c7bb-dad298ce5ded", false},

		{"a84dfca-43e0-46ea-a7bb-dad298ce5ded", false},
		{"4a84dfca-3e0-46ea-a7bb-dad298ce5ded", false},
		{"4a84dfca-43e0-6ea-a7bb-dad298ce5ded", false},
		{"4a84dfca-43e0-46ea-7bb-dad298ce5ded", false},
		{"4a84dfca-43e0-46ea-a7bb-ad298ce5ded", false},

		{" 4a84dfca-43e0-46ea-a7bb-dad298ce5ded", false},
		{"4a84dfca-43e0-46ea -a7bb-dad298ce5ded", false},
		{"4a84dfca-43e0-46ea-a7bb-dad298ce5ded ", false},

		{"ga84dfca-43e0-46ea-a7bb-dad298ce5ded", false},

		{"random", false},
		{" ", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsValidUUIDv4{Name: "UUIDv4", Field: test.field}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsValidUUIDv4Error(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
