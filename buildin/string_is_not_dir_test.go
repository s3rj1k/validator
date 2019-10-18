package buildin

import (
	"fmt"
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsNotDir(t *testing.T) {
	r := require.New(t)

	fd, err := os.Create(regularFile)
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	v := StringIsNotDir{Name: "Name", Field: regularFile}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	err = os.Remove(regularFile)
	r.Nil(err)

	v = StringIsNotDir{Name: "Name", Field: regularFile}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	err = os.MkdirAll(dir, 0666)
	r.Nil(err)

	v = StringIsNotDir{Name: "Name", Field: dir}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{fmt.Sprintf("path '%s' is a dir", dir)}, e.Get("Name"))

	err = os.RemoveAll(dir)
	r.Nil(err)

	v = StringIsNotDir{Name: "Name", Field: dir}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())
}
