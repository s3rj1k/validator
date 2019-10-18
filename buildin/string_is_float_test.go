package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsFloat(t *testing.T) {
	r := require.New(t)

	v := &StringIsFloat{Name: "Name", Field: "15.22"} // must be float
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsFloat{Name: "Name", Field: "-0.0"} // must be float
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsFloat{Name: "Name", Field: ""} // empty string is valid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsFloat{Name: "Name", Field: "135"} // int is valid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsFloat{Name: "Name", Field: "baby"} // string is invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsFloatError(v)}, e.Get("Name"))

	v = &StringIsFloat{Name: "Name", Field: "123.2 "} // whitespaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsFloatError(v)}, e.Get("Name"))

	v = &StringIsFloat{Name: "Name", Field: "   "} // only whitespaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsFloatError(v)}, e.Get("Name"))
}
