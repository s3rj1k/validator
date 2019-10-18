package buildin

import (
	"fmt"
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsNotPath(t *testing.T) {
	r := require.New(t)

	fd, err := os.Create(regularFile) // nolint: gosec
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	defer func(r *require.Assertions) {
		err = os.Remove(regularFile)
		r.Nil(err)
	}(r)

	v := StringIsNotPath{Name: "Name", Field: notExists}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsNotPath{Name: "Name", Field: regularFile}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{fmt.Sprintf("path '%s' must not exist", regularFile)}, e.Get("Name"))

	v = StringIsNotPath{Name: "Name", Field: regularFile}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{fmt.Sprintf("path '%s' must not exist", regularFile)}, e.Get("Name"))
}
