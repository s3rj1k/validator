package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsRGBcolor(t *testing.T) {

	r := require.New(t)

	v := &StringIsRGBcolor{Name: "Name", Field: "rgb(0,0,0)"} // hexcolor here
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsRGBcolor{Name: "Name", Field: "rgb(255,255,255)"} // hexcolor here also (3-6 0-F chars)
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsRGBcolor{Name: "Name", Field: "RGB(0,15,25)"} // rgb must be lowercased
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsRGBcolorError(v)}, e.Get("Name"))

	v = &StringIsRGBcolor{Name: "Name", Field: "rgb(0,0,256)"} // values 0-255
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsRGBcolorError(v)}, e.Get("Name"))

	v = &StringIsRGBcolor{Name: "Name", Field: "    "} // empty string or only whitespaces are invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsRGBcolorError(v)}, e.Get("Name"))

	v = &StringIsRGBcolor{Name: "Name", Field: "ffd762 "} // whitespaces are not trimmed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsRGBcolorError(v)}, e.Get("Name"))
}
