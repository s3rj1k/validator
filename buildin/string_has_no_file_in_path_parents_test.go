package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringHasNoFilesInPathParents(t *testing.T) {
	r := require.New(t)

	v := StringHasNoFileInPathParents{Name: "Name", Field: "."} // at least 1 upper case
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringHasNoFileInPathParents{Name: "Name", Field: "../Makefile"} // at least 1 upper case
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
}
