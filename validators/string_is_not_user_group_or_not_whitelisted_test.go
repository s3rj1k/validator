package validators

import (
	"os/user"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsNotUserGroupOrNotWhitelisted(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field     string
		whitelist []string
		valid     bool
	}{
		{"", []string{""}, false},
		{"baby", []string{"baby"}, false},
		{"baby", nil, true},

		{" ", nil, true},
		{"", nil, true},
	}

	cu, err := user.Current()
	if err == nil && cu.Gid != "0" {
		tests[0].field = cu.Name
	}

	for index, test := range tests {
		v := &StringIsNotUserGroupOrNotWhitelisted{Name: "Passwd", Field: test.field, Whitelist: test.whitelist}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%s got=%s", index, !test.valid, e.HasAny())
		if !test.valid {
			r.Equalf([]string{StringIsNotUserGroupOrNotWhitelistedError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
