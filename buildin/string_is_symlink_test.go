package buildin

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsSymlink(t *testing.T) {

	r := require.New(t)

	fd, err := os.Create("/tmp/test")
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	v := StringIsSymlink{Name: "Name", Field: "/tmp/test"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())

	err = os.Remove("/tmp/test")
	r.Nil(err)

	v = StringIsSymlink{Name: "Name", Field: "/tmp/test"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())

	_ = os.Remove("/tmp/test_symlink")
	err = os.Symlink("/tmp", "/tmp/test_symlink")
	r.Nil(err)

	v = StringIsSymlink{Name: "Name", Field: "/tmp/test_symlink"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	err = os.Remove("/tmp/test_symlink")
	r.Nil(err)

	v = StringIsSymlink{Name: "Name", Field: "/tmp/test_symlink"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path '/tmp/test_symlink' is not a symlink"}, e.Get("Name"))
}
