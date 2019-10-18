package buildin

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsPathAndNotSocket(t *testing.T) {
	r := require.New(t)

	// setup
	_silentdeleteMany(notExists, regularFile, unixSocket, fileWithSetUID, fileWithSetGID)
	_createfiles(regularFile, fileWithSetGID, fileWithSetUID)
	_setfilemode(fileWithSetGID, os.ModeSetgid)
	_setfilemode(fileWithSetUID, os.ModeSetuid)
	_creatunixsocket(unixSocket)

	// teardown
	defer _silentdeleteMany(regularFile, unixSocket, fileWithSetUID, fileWithSetGID)

	var tests = []struct {
		field string
		valid bool
	}{
		{notExists, false},

		{regularFile, true},

		{unixSocket, false},

		{fileWithSetUID, true},
		{fileWithSetGID, true},

		{"", false},
	}

	for index, test := range tests {
		v := &StringIsPathAndNotSocket{Name: "FileModes", Field: test.field}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsPathAndNotSocketError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
