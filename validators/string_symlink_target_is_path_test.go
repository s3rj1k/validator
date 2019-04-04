package validators

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringSymlinkTargetIsPath(t *testing.T) {

	r := require.New(t)

	fd, err := os.Create("/tmp/not_a_symlink") // nolint: gosec
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	v := &StringSymlinkTargetIsPath{Name: "Name", Field: "/tmp/not_a_symlink"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count()) // no error, assess only symlinks

	_ = os.Remove("/tmp/test_symlink")
	err = os.Symlink("/tmp", "/tmp/test_symlink") // symlink to folder
	r.Nil(err)

	v = &StringSymlinkTargetIsPath{Name: "Name", Field: "/tmp/test_symlink"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count()) // symlink to folder is OK

	_ = os.Remove("/tmp/test_symlink")
	err = os.Symlink("/tmp/not_a_symlink", "/tmp/test_symlink") // symlink to file
	r.Nil(err)

	v = &StringSymlinkTargetIsPath{Name: "Name", Field: "/tmp/test_symlink"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count()) // symlink to file is OK

	err = os.Remove("/tmp/not_a_symlink")
	r.Nil(err)

	v = &StringSymlinkTargetIsPath{Name: "Name", Field: "/tmp/test_symlink"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count()) // target does not exist is error
	r.Equal([]string{StringSymlinkTargetIsPathError(v)}, e.Get("Name"))

	err = os.Remove("/tmp/test_symlink")
	r.Nil(err)
}