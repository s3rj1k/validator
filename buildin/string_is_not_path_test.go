package buildin

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsNotPath(t *testing.T) {

	r := require.New(t)

	fd, err := os.Create("/tmp/test") // nolint: gosec
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	defer func(r *require.Assertions) {
		err = os.Remove("/tmp/test")
		r.Nil(err)
	}(r)

	v := StringIsNotPath{Name: "Name", Field: "/tmp/doesnotexist"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsNotPath{Name: "Name", Field: "/tmp/test"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path '/tmp/test' must not exist"}, e.Get("Name"))

	v = StringIsNotPath{Name: "Name", Field: "/tmp/test"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path '/tmp/test' must not exist"}, e.Get("Name"))
}