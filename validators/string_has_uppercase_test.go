package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringHasUpperCase(t *testing.T) {

	r := require.New(t)

	v := &StringHasUpperCase{Name: "Name", Field: "3w4asF@`"} // at least 1 upper case
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringHasUpperCase{Name: "Name", Field: ""} // empty string is valid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringHasUpperCase{Name: "Name", Field: " Space Inside "} // outer and inner whitespaces are allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringHasUpperCase{Name: "Name", Field: "abc123"} // must contain uppercase
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringHasUpperCaseError(v)}, e.Get("Name"))

	v = &StringHasUpperCase{Name: "Name", Field: "    "} // only spaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringHasUpperCaseError(v)}, e.Get("Name"))
}
