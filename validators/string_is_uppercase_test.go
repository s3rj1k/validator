package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsUpperCase(t *testing.T) {

	r := require.New(t)

	v := StringIsUpperCase{Name: "Name", Field: "ASFADG44"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsUpperCase{Name: "Name", Field: "   "} // empty string is valid, spaces are trimmed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsUpperCase{Name: "Name", Field: " A5555 "} // outer whitespaces are allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsUpperCase{Name: "Name", Field: "AD GGGG"} // inner whitespaces are allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsUpperCase{Name: "Name", Field: "Abcd"} // lowercase is invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be uppercase"}, e.Get("Name"))
}
