package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringMatchRegex(t *testing.T) {

	r := require.New(t)

	v := StringMatchRegex{Name: "Phone", Field: "555-555-5555", Regex: "^([0-9]{3}-[0-9]{3}-[0-9]{4})$"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringMatchRegex{Name: "Phone", Field: "123-ab1-1424", Regex: "^([0-9]{3}-[0-9]{3}-[0-9]{4})$"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Phone does not match the expected format"}, e.Get("Phone"))

	v = StringMatchRegex{"Phone", "123-ab1-1424", "^([0-9]{3}-[0-9]{3}-[0-9]{4})$"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Phone does not match the expected format"}, e.Get("Phone"))
}
