package buildin

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsHexadecimal(t *testing.T) {
	r := require.New(t)

	v := &StringIsHexadecimal{Name: "Name", Field: hex.EncodeToString([]byte("Hello"))} // hex here
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsHexadecimal{Name: "Name", Field: "FF"} // hex here also
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsHexadecimal{Name: "Name", Field: fmt.Sprintf("%x", 155)} // and this is hex
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsHexadecimal{Name: "Name", Field: "Hello"} // other strings are invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsHexadecimalError(v)}, e.Get("Name"))

	v = &StringIsHexadecimal{Name: "Name", Field: ""} // empty string is invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsHexadecimalError(v)}, e.Get("Name"))

	v = &StringIsHexadecimal{Name: "Name", Field: "    "} // whitespaces are invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsHexadecimalError(v)}, e.Get("Name"))

	v = &StringIsHexadecimal{Name: "Name", Field: "FF "} // whitespaces are not trimmed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsHexadecimalError(v)}, e.Get("Name"))
}
