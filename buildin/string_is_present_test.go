package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsPresent(t *testing.T) {

	r := require.New(t)

	v := &StringIsPresent{Name: "Name", Field: "Mark"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsPresent{Name: "Name", Field: " "}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsPresentError(v)}, e.Get("Name"))

	v = &StringIsPresent{Name: "Name", Field: ""}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsPresentError(v)}, e.Get("Name"))
}
