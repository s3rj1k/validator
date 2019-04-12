package buildin

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsDir(t *testing.T) {

	r := require.New(t)

	fd, err := os.Create("/tmp/test")
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	v := StringIsDir{Name: "Name", Field: "/tmp/test"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())

	err = os.Remove("/tmp/test")
	r.Nil(err)

	v = StringIsDir{Name: "Name", Field: "/tmp/test"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())

	err = os.MkdirAll("/tmp/test_dir", 0666)
	r.Nil(err)

	v = StringIsDir{Name: "Name", Field: "/tmp/test_dir"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	err = os.Remove("/tmp/test_dir")
	r.Nil(err)

	v = StringIsDir{Name: "Name", Field: "/tmp/test_dir"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path '/tmp/test_dir' is not a dir"}, e.Get("Name"))
}
