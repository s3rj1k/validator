package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsUTFLetterNum(t *testing.T) {

	r := require.New(t)

	v := &StringIsUTFLetterNum{Name: "Name", Field: "a১522௫sd品ʂля٣"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsUTFLetterNum{Name: "Name", Field: ""} // empty string is valid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsUTFLetterNum{Name: "Name", Field: ":~$"} // any other characters except for UTF letters are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsUTFLetterNumError(v)}, e.Get("Name"))

	v = &StringIsUTFLetterNum{Name: "Name", Field: " ля 品ʂ "} // inner/outer whitespaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsUTFLetterNumError(v)}, e.Get("Name"))

	v = &StringIsUTFLetterNum{Name: "Name", Field: "   "} // only whitespaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsUTFLetterNumError(v)}, e.Get("Name"))
}
