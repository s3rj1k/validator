package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsNull(t *testing.T) {
	r := require.New(t)

	v := &StringIsNull{Name: "Name", Field: ""}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsNull{Name: "Name", Field: *new(string)}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsNull{Name: "Name", Field: " "} // whitespaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsNullError(v)}, e.Get("Name"))
}
