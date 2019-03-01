package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsJSONDive(t *testing.T) {

	r := require.New(t)

	field := []string{"{\"test\": \"sure\"}", "   {\"test\": \"sure\"}    ", "123"}

	v := StringSliceDive{
		Validator: &StringIsJSON{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{" ", "abc", ""}

	v = StringSliceDive{
		Validator: &StringIsJSON{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(3, e.Count())
}
