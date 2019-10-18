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

	_ = os.Remove(writableFile)

	fd, err := os.Create(writableFile)
	r.Nil(err)

	v := &StringIsPathAndIsWritable{Name: "Name", Field: writableFile}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	err = fd.Chmod(0000)
	r.Nil(err)

	v = &StringIsPathAndIsWritable{Name: "Name", Field: writableFile}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsPathAndIsWritableError(v)}, e.Get("Name"))

	err = fd.Chmod(0777)
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	err = os.Remove(writableFile)
	r.Nil(err)

	_ = os.Remove(writableDir)

	err = os.MkdirAll(writableDir, 0777)
	r.Nil(err)

	v = &StringIsPathAndIsWritable{Name: "Name", Field: writableDir}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	err = os.Chmod(writableDir, 0000)
	r.Nil(err)

	v = &StringIsPathAndIsWritable{Name: "Name", Field: writableDir}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsPathAndIsWritableError(v)}, e.Get("Name"))

	err = os.Chmod(writableDir, 0777)
	r.Nil(err)

	err = os.Remove(writableDir)
	r.Nil(err)
}
