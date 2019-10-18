package buildin

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsDirWithModeSticky(t *testing.T) {
	r := require.New(t)

	// setup
	_silentdeleteMany(notExists, notDir, notStickyDir, stickyDir)
	_createdirs(notStickyDir, stickyDir)
	_setfilemode(stickyDir, os.ModeSticky)
	_createfile(notDir)

	// teardown
	defer _silentdeleteMany(notStickyDir, stickyDir, notDir)

	var tests = []struct {
		field string
		valid bool
	}{
		{stickyDir, true},

		{notExists, false},
		{notDir, false},
		{notStickyDir, false},

		{"", false},
	}

	for index, test := range tests {
		v := &StringIsDirWithModeSticky{Name: "StickyDir", Field: test.field}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsDirWithModeStickyError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
