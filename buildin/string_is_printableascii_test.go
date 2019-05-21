package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsPrintableASCII(t *testing.T) {
	r := require.New(t)

	v := &StringIsPrintableASCII{Name: "Name", Field: "abc123"} // must be ASCII
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsPrintableASCII{Name: "Name", Field: "!$#%()-=<>etc...,@"} // must be ASCII
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsPrintableASCII{Name: "Name", Field: ""} // empty string is valid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsPrintableASCII{Name: "Name", Field: " 123 "} // outer whitespaces are allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsPrintableASCII{Name: "Name", Field: "   "} // only whitespaces are allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsPrintableASCII{Name: "Name", Field: string(rune(10))} // non-printable in invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsPrintableASCIIError(v)}, e.Get("Name"))

	v = &StringIsPrintableASCII{Name: "Name", Field: "опа"} // non-ascii in invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsPrintableASCIIError(v)}, e.Get("Name"))
}
