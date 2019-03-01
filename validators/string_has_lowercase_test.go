package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringHasLowerCase(t *testing.T) {

	r := require.New(t)

	v := StringHasLowerCase{Name: "Name", Field: "3w4ASF^^#"} // at least 1 lowercase
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringHasLowerCase{Name: "Name", Field: ""} // empty string is valid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringHasLowerCase{Name: "Name", Field: " space inside "} // outer and inner whitespaces are allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringHasLowerCase{Name: "Name", Field: "ABC123"} // must contain lowercase
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain at least 1 lowercase"}, e.Get("Name"))

	v = StringHasLowerCase{Name: "Name", Field: "    "} // only spaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must contain at least 1 lowercase"}, e.Get("Name"))
}
