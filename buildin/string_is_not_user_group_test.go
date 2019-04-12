package buildin

import (
	"os/user"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsNotUserGroup(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"not_user_group", true},

		{" ", true},
		{"", true},
	}

	// trying to add current user in test if not root
	cu, err := user.Current()
	if err == nil && cu.Gid != "0" {
		tests = append(tests, struct {
			field string
			valid bool
		}{cu.Username, false})
	}

	for index, test := range tests {
		v := &StringIsNotUserGroup{Name: "Passwd", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())
		if !test.valid {
			r.Equalf([]string{StringIsNotUserGroupError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
