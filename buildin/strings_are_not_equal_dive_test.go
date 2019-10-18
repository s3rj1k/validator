package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringAreNotEqualDive(t *testing.T) {
	r := require.New(t)

	sl := []string{"Foo", "Bar", "Bob", "What?"}

	v := StringSliceDive{
		Validator: &StringsAreEqual{Name: "Slice", ComparedField: "Bar"},
		Field:     sl,
	}
	e := validator.NewErrorsP()

	v.Validate(e)
	r.Equal(3, e.Count())

	v = StringSliceDive{
		Validator: &StringsAreEqual{Name: "Slice", ComparedField: "Bar"},
		Field:     []string{"Bar", "Bar", "bar", ""},
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(2, e.Count())
}
