package validators

import (
	"encoding/base32"
	"encoding/base64"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsBase64(t *testing.T) {

	r := require.New(t)

	sEnc := base64.StdEncoding.EncodeToString([]byte("abc123"))
	v := StringIsBase64{Name: "Name", Field: sEnc} // must be base64 string
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringIsBase64{Name: "Name", Field: ""} // empty string is invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be base64 encoded"}, e.Get("Name"))

	v = StringIsBase64{Name: "Name", Field: " " + sEnc + " "} // outer whitespaces are invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be base64 encoded"}, e.Get("Name"))

	v = StringIsBase64{Name: "Name", Field: "   "} // only whitespaces are invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be base64 encoded"}, e.Get("Name"))

	sEnc32 := base32.StdEncoding.EncodeToString([]byte("abc123"))
	v = StringIsBase64{Name: "Name", Field: sEnc32} // base32 is bad
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be base64 encoded"}, e.Get("Name"))

	v = StringIsBase64{Name: "Name", Field: "abc123"} // simple string is bad
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name must be base64 encoded"}, e.Get("Name"))
}
