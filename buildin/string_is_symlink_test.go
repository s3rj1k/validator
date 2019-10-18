package buildin

import (
	"fmt"
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsSymlink(t *testing.T) {
	r := require.New(t)

	fd, err := os.Create(regularFile)
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	v := StringIsSymlink{Name: "Name", Field: regularFile}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(1, e.Count())

	err = os.Remove(regularFile)
	r.Nil(err)

	v = StringIsSymlink{Name: "Name", Field: regularFile}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())

	_ = os.Remove(symlink)
	err = os.Symlink("/tmp", symlink)
	r.Nil(err)

	v = StringIsSymlink{Name: "Name", Field: symlink}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	err = os.Remove(symlink)
	r.Nil(err)

	v = StringIsSymlink{Name: "Name", Field: symlink}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{fmt.Sprintf("path '%s' is not a symlink", symlink)}, e.Get("Name"))
}
