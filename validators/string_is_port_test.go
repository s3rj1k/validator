package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsPort(t *testing.T) {

	r := require.New(t)

	v := StringIsPort{Name: "Name", Field: "1"} // Port is OK > 0
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsPort{Name: "Name", Field: "65535"} // Port is OK < 65536
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsPort{Name: "Name", Field: ""} // empty string is invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must represent a valid port"}, e.Get("Name"))

	v = StringIsPort{Name: "Name", Field: " 13 "} // outer whitespaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must represent a valid port"}, e.Get("Name"))
}
