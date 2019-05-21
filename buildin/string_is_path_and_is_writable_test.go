package buildin

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsPathAndIsWritable(t *testing.T) {
	r := require.New(t)

	_ = os.Remove("/tmp/string_writable_file")

	fd, err := os.Create("/tmp/string_writable_file")
	r.Nil(err)

	v := &StringIsPathAndIsWritable{Name: "Name", Field: "/tmp/string_writable_file"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	err = fd.Chmod(0000)
	r.Nil(err)

	v = &StringIsPathAndIsWritable{Name: "Name", Field: "/tmp/string_writable_file"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsPathAndIsWritableError(v)}, e.Get("Name"))

	err = fd.Chmod(0777)
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	err = os.Remove("/tmp/string_writable_file")
	r.Nil(err)

	_ = os.Remove("/tmp/string_writable_dir")

	err = os.MkdirAll("/tmp/string_writable_dir", 0777)
	r.Nil(err)

	v = &StringIsPathAndIsWritable{Name: "Name", Field: "/tmp/string_writable_dir"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	err = os.Chmod("/tmp/string_writable_dir", 0000)
	r.Nil(err)

	v = &StringIsPathAndIsWritable{Name: "Name", Field: "/tmp/string_writable_dir"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsPathAndIsWritableError(v)}, e.Get("Name"))

	err = os.Chmod("/tmp/string_writable_dir", 0777)
	r.Nil(err)

	err = os.Remove("/tmp/string_writable_dir")
	r.Nil(err)
}
