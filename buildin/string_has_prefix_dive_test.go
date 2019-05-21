package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringHasPrefixDive(t *testing.T) {
	r := require.New(t)

	sl := []string{"Foo", "Bar", "Bob", "What?"}

	v := StringSliceDive{
		Validator: &StringHasPrefix{Name: "Slice", ComparedField: "W"},
		Field:     sl,
	}
	e := validator.NewErrorsP()
	v.Validate(e)
	r.Equal(3, e.Count()) // 4(total) - 1 strings in sl that have matched prefix
}
