package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsInt(t *testing.T) {
	r := require.New(t)

	v := &StringIsInt{Name: "Name", Field: "1988"} // must be int
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsInt{Name: "Name", Field: "-0"} // must be int
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsInt{Name: "Name", Field: ""} // empty string is valid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsInt{Name: "Name", Field: "13.5"} // float is invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsIntError(v)}, e.Get("Name"))

	v = &StringIsInt{Name: "Name", Field: "baby"} // string is invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsIntError(v)}, e.Get("Name"))

	v = &StringIsInt{Name: "Name", Field: "123 "} // whitespaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsIntError(v)}, e.Get("Name"))

	v = &StringIsInt{Name: "Name", Field: "   "} // only whitespaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsIntError(v)}, e.Get("Name"))
}
