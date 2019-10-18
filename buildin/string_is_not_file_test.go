package buildin

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsNotFile(t *testing.T) {
	r := require.New(t)

	fd, err := os.Create(regularFile)
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	v := &StringIsNotFile{Name: "Name", Field: regularFile}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsNotFileError(v)}, e.Get("Name"))

	err = os.Remove(regularFile)
	r.Nil(err)

	v = &StringIsNotFile{Name: "Name", Field: regularFile}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsNotFile{Name: "Name", Field: "/tmp"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())
}
