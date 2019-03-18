package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsJSON(t *testing.T) {

	r := require.New(t)

	v := &StringIsJSON{Name: "Name", Field: "{\"test\": \"sure\"}"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsJSON{Name: "Name", Field: "   {\"test\": \"sure\"}    "} // outer whitespaces are valid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsJSON{Name: "Name", Field: ""} // empty string is invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsJSONError(v)}, e.Get("Name"))
}
