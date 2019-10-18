package buildin

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsSymlinkAndTargetIsPath(t *testing.T) {
	r := require.New(t)

	fd, err := os.Create(notSymlink) // nolint: gosec
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	v := &StringIsSymlinkAndTargetIsPath{Name: "Name", Field: notSymlink}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count()) // no error, assess only symlinks

	_ = os.Remove(symlink)
	err = os.Symlink("/tmp", symlink) // symlink to folder
	r.Nil(err)

	v = &StringIsSymlinkAndTargetIsPath{Name: "Name", Field: symlink}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count()) // symlink to folder is OK

	_ = os.Remove(symlink)
	err = os.Symlink(notSymlink, symlink) // symlink to file
	r.Nil(err)

	v = &StringIsSymlinkAndTargetIsPath{Name: "Name", Field: symlink}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count()) // symlink to file is OK

	err = os.Remove(notSymlink)
	r.Nil(err)

	v = &StringIsSymlinkAndTargetIsPath{Name: "Name", Field: symlink}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count()) // target does not exist is error
	r.Equal([]string{StringIsSymlinkAndTargetIsPathError(v)}, e.Get("Name"))

	err = os.Remove(symlink)
	r.Nil(err)
}
