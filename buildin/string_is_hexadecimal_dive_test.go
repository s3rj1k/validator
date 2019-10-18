package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsHexadecimalDive(t *testing.T) {
	r := require.New(t)

	field := []string{"123456789", "abcdef", "ABCDEF"}

	v := StringSliceDive{
		Validator: &StringIsHexadecimal{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"123456789", "abcdef", "ABCDEF", " a", "", "Z", "  "}

	v = StringSliceDive{
		Validator: &StringIsHexadecimal{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(4, e.Count())
}
