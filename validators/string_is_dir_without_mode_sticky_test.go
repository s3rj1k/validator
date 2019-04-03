package validators

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsDirWithoutModeSticky(t *testing.T) {

	r := require.New(t)

	var notexists, notdir, notstickydir, stickydir string
	notexists = "/tmp/testnotexists"
	notdir = "/tmp/testnotdir"
	notstickydir = "/tmp/testnotstickydir"
	stickydir = "/tmp/teststickydir"

	// setup
	_silentdeleteMany(notexists, notdir, notstickydir, stickydir)
	_createdirs(notstickydir, stickydir)
	_setfilemode(stickydir, os.ModeSticky)
	_createfile(notdir)

	// teardown
	defer _silentdeleteMany(notstickydir, stickydir, notdir)

	var tests = []struct {
		field string
		valid bool
	}{
		{notstickydir, true},

		{notexists, false},
		{notdir, false},
		{stickydir, false},

		{"", false},
	}

	for index, test := range tests {
		v := &StringIsDirWithoutModeSticky{Name: "StickyDir", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsDirWithoutModeStickyError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
