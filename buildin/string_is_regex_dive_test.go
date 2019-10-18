package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsRegexDive(t *testing.T) {
	r := require.New(t)

	field := []string{Email, Float, Integer}

	v := StringSliceDive{
		Validator: &StringIsRegex{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{Email, "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9+))?$", Integer}

	v = StringSliceDive{
		Validator: &StringIsRegex{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
}
