package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsMAC(t *testing.T) {

	r := require.New(t)

	v := StringIsMAC{Name: "Name", Field: "01:23:45:67:89:ab"} // MAC is OK
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsMAC{Name: "Name", Field: ""} // empty string is invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be valid MAC address"}, e.Get("Name"))

	v = StringIsMAC{Name: "Name", Field: " 01:23:45:67:89:ab "} // outer whitespaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be valid MAC address"}, e.Get("Name"))
}
