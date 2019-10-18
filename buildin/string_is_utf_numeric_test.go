package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsUTFNumeric(t *testing.T) {
	r := require.New(t)

	v := &StringIsUTFNumeric{Name: "Name", Field: "১522௫٣"}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsUTFNumeric{Name: "Name", Field: ""} // empty string is valid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsUTFNumeric{Name: "Name", Field: "ag:~$"} // any other characters except for UTF letters are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsUTFNumericError(v)}, e.Get("Name"))

	v = &StringIsUTFNumeric{Name: "Name", Field: " ля 品ʂ "} // inner/outer whitespaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsUTFNumericError(v)}, e.Get("Name"))

	v = &StringIsUTFNumeric{Name: "Name", Field: "   "} // only whitespaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsUTFNumericError(v)}, e.Get("Name"))
}
