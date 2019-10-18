package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsAbsPath(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{regularFile + "/", true},
		{regularFile, true},
		{"/tmp", true},
		{"/", true},
		{"/tmp/ test/", true},
		{"/tmp/test ", true},
		{"/tmp/test/ ", true},
		{"/tmp/test/ /", true},

		{"/tmp/test//", false},
		{"/tmp/test///", false},
		{"/tmp/test/ //", false},
		{"//", false},
		{"test", false},
		{"/tmp//test/test", false},
		{"./test", false},

		{" ", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsAbsPath{Name: "AbsPath", Field: test.field}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsAbsPathError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
