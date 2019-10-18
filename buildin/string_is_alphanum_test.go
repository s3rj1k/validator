package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsAlphaNum(t *testing.T) {
	r := require.New(t)

	v := &StringIsAlphaNum{Name: "Name", Field: "ASfgg5452"}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsAlphaNum{Name: "Name", Field: ""} // empty string is valid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsAlphaNum{Name: "Name", Field: "ы$^"} // any other characters except for a-zA-Z are invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsAlphaNumError(v)}, e.Get("Name"))

	v = &StringIsAlphaNum{Name: "Name", Field: " wh1t3 spaces "} // inner/outer whitespaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsAlphaNumError(v)}, e.Get("Name"))

	v = &StringIsAlphaNum{Name: "Name", Field: "   "} // only whitespaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsAlphaNumError(v)}, e.Get("Name"))
}
