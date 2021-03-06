package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsLowerCaseDive(t *testing.T) {
	r := require.New(t)

	field := []string{"each", " is ", "lower case", "", " "} // 0 errors

	v := StringSliceDive{
		Validator: &StringIsLowerCase{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"TWO", "UPPERCASED", "oh", "no", "Three", ""} // 2 errors

	v = StringSliceDive{
		Validator: &StringIsLowerCase{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(3, e.Count())
}
