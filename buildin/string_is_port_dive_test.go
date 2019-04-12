package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsPortDive(t *testing.T) {

	r := require.New(t)

	field := []string{"1", "123", "65535"}

	v := StringSliceDive{
		Validator: &StringIsPort{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"1", "123", "65535", "1 ", "65536", " ", ""}

	v = StringSliceDive{
		Validator: &StringIsPort{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(4, e.Count())
}
