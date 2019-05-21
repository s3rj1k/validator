package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsASCII(t *testing.T) {
	r := require.New(t)

	v := &StringIsASCII{Name: "Name", Field: "abc123"} // must be ASCII
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsASCII{Name: "Name", Field: "!$#%()-=<>etc...,@"} // must be ASCII
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsASCII{Name: "Name", Field: ""} // empty string is valid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsASCII{Name: "Name", Field: " 123 "} // outer whitespaces are allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsASCII{Name: "Name", Field: "   "} // only whitespaces are allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsASCII{Name: "Name", Field: "опа"} // non-ascii in invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsASCIIError(v)}, e.Get("Name"))
}
