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

	fd, err := os.Create("/tmp/string_is_file_test")
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	v := &StringIsNotFile{Name: "Name", Field: "/tmp/string_is_file_test"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsNotFileError(v)}, e.Get("Name"))

	err = os.Remove("/tmp/string_is_file_test")
	r.Nil(err)

	v = &StringIsNotFile{Name: "Name", Field: "/tmp/string_is_file_test"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsNotFile{Name: "Name", Field: "/tmp"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())
}
