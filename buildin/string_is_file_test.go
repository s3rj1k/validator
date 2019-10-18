package buildin

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsFile(t *testing.T) {
	r := require.New(t)

	fd, err := os.Create(regularFile)
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	v := &StringIsFile{Name: "Name", Field: regularFile}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	err = os.Remove(regularFile)
	r.Nil(err)

	v = &StringIsFile{Name: "Name", Field: regularFile}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsFileError(v)}, e.Get("Name"))

	v = &StringIsFile{Name: "Name", Field: "/tmp"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsFileError(v)}, e.Get("Name"))
}
