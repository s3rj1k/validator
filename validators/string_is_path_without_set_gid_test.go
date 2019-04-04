package validators

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsPathWithoutSetGID(t *testing.T) {

	r := require.New(t)

	var notexists, regularFile, unixSocket, fileWithSetUID, fileWithSetGID string
	notexists = "/tmp/testnotexists"
	regularFile = "/tmp/testregularFile"
	unixSocket = "/tmp/testunixSocket"
	fileWithSetUID = "/tmp/testfileWithSetUID"
	fileWithSetGID = "/tmp/testfileWithSetGID"

	// setup
	_silentdeleteMany(notexists, regularFile, unixSocket, fileWithSetUID, fileWithSetGID)
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
		{notexists, false},

		{regularFile, true},

		{unixSocket, true},

		{fileWithSetUID, true},
		{fileWithSetGID, false},

		{"", false},
	}

	for index, test := range tests {
		v := &StringIsPathWithoutSetGID{Name: "FileModes", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsPathWithoutSetGIDError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
