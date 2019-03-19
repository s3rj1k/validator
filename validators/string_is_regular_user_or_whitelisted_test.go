package validators

import (
	"os/user"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsRegularUserOrWhitelisted(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field     string
		whitelist []string
		valid     bool
	}{
		{"", []string{""}, true},
		{"baby", []string{"baby"}, true},
		{"baby", nil, false},

		{" ", nil, false},
		{"", nil, false},
	}

	cu, err := user.Current()
	if err == nil && cu.Gid != "0" {
		tests[0].field = cu.Name
	}

	for index, test := range tests {
		v := &StringIsRegularUserOrWhitelisted{Name: "Passwd", Field: test.field, Whitelist: test.whitelist}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())
		
		if !test.valid {
			r.Equalf([]string{StringIsRegularUserOrWhitelistedError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
