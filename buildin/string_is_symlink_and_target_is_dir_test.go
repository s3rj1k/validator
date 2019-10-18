package buildin

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsSymlinkAndTargetIsDir(t *testing.T) {
	r := require.New(t)

	fd, err := os.Create(notSymlink) // nolint: gosec
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	v := &StringIsSymlinkAndTargetIsDir{Name: "Name", Field: notSymlink}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count()) // not a symlink is OK

	_ = os.Remove(symlink)
	err = os.Symlink(notSymlink, symlink) // symlink to file
	r.Nil(err)

	v = &StringIsSymlinkAndTargetIsDir{Name: "Name", Field: symlink}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count()) // symlink to file is error
	r.Equal([]string{StringIsSymlinkAndTargetIsDirError(v)}, e.Get("Name"))

	_ = os.Remove(symlink)
	err = os.Symlink("/tmp", symlink) // symlink to folder
	r.Nil(err)

	v = &StringIsSymlinkAndTargetIsDir{Name: "Name", Field: symlink}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count()) // symlink to folder is good

	err = os.Remove(symlink)
	r.Nil(err)
	err = os.Remove(notSymlink)
	r.Nil(err)

	v = &StringIsSymlinkAndTargetIsDir{Name: "Name", Field: symlink}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count()) // does not exist is OK
}
