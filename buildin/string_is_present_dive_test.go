package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsPresentDive(t *testing.T) {
	r := require.New(t)

	field := []string{"abc", " a", "3  3"}

	v := StringSliceDive{
		Validator: &StringIsPresent{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"", "   ", "1", "a"}

	v = StringSliceDive{
		Validator: &StringIsPresent{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(2, e.Count())
}
