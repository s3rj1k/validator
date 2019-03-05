package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsBase64Dive(t *testing.T) {

	r := require.New(t)

	field := []string{"xoP4nZV8Gv9ceg==", "n/GNBg=="} // 0 errors

	v := StringSliceDive{
		Validator: &StringIsBase64{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"xoP4nZV8Gv9ceg==", "n/GNBg==", "xoP4nZ V8Gv9ceg==", " n/GNBg==", " ", ""} // 3 errors

	v = StringSliceDive{
		Validator: &StringIsBase64{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(3, e.Count())
}
