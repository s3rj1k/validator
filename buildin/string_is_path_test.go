package buildin

import (
	"fmt"
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsPath(t *testing.T) {
	r := require.New(t)

	fd, err := os.Create(regularFile) // nolint: gosec
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	defer func(r *require.Assertions) {
		err = os.Remove(regularFile)
		r.Nil(err)
	}(r)

	v := StringIsPath{Name: "Name", Field: regularFile}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsPath{Name: "Name", Field: notExists}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{fmt.Sprintf("path '%s' must exist", notExists)}, e.Get("Name"))

	e = validator.NewErrors()
	v = StringIsPath{Name: "Name", Field: ""}
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"path '' must exist"}, e.Get("Name"))
}
