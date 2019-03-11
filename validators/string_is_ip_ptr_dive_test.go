package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsIPptrDive(t *testing.T) {

	r := require.New(t)

	field := []string{"8.8.8.8", "127.0.0.1", "2001:4860:4860::8888"}

	v := StringSliceDive{
		Validator: &StringIsIPptr{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"265.255.253.0", "127.0.0.", "0:0:0:0:0:0", " 127.0.0.1", "127.0. 0.1", " ", ""}

	v = StringSliceDive{
		Validator: &StringIsIPptr{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(7, e.Count())
}
