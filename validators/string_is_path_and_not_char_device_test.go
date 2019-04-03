package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsPathAndNotCharDevice(t *testing.T) {

	r := require.New(t)

	var notexists, namedPipe string
	notexists = "/tmp/testnotexists"
	namedPipe = "/tmp/testnamedPipe"

	// setup
	_silentdeleteMany(notexists, namedPipe)
	_createnamedpipe(namedPipe)

	// teardown
	defer _silentdeleteMany(namedPipe)

	var tests = []struct {
		field string
		valid bool
	}{
		{"/dev/tty", false},

		{notexists, false},
		{namedPipe, true},
		{"/dev/null", false},

		{"", false},
	}

	for index, test := range tests {
		v := &StringIsPathAndNotCharDevice{Name: "FileModes", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsPathAndNotCharDeviceError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
