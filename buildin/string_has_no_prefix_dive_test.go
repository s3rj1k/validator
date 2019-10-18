package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringHasNoPrefixDive(t *testing.T) {
	r := require.New(t)

	sl := []string{"Foo", "Bar", "Bob", "What?"}

	v := StringSliceDive{
		Validator: &StringHasNoPrefix{Name: "Slice", ComparedField: "F"},
		Field:     sl,
	}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(1, e.Count()) // 4(total) - 3 strings in sl that do not have matched prefix
}
