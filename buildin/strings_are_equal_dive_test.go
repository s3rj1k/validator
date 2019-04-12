package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringAreEqualDive(t *testing.T) {

	r := require.New(t)

	sl := []string{"Foo", "Bar", "Bob", "What?"}

	v := StringSliceDive{
		Validator: &StringsAreEqual{Name: "Slice", ComparedField: "Bar"},
		Field:     sl,
	}
	e := validator.NewErrorsP()
	v.Validate(e)
	r.Equal(3, e.Count()) // 3 strings in sl that do not match compared

	v = StringSliceDive{
		Validator: &StringsAreEqual{Name: "Slice", ComparedField: "Bar"},
		Field:     []string{"Bar", "Bar", "bar", ""},
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(2, e.Count()) // empty string and lowercased
}
