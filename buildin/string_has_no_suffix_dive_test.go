package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringHasNoSuffixDive(t *testing.T) {
	r := require.New(t)

	sl := []string{"Foo", "Bar", "Bob", "What?"}

	v := StringSliceDive{
		Validator: &StringHasNoSuffix{Name: "Slice", ComparedField: "r"},
		Field:     sl,
	}
	e := validator.NewErrorsP()
	v.Validate(e)
	r.Equal(1, e.Count()) // 4(total) - 3 strings in sl that do not have matched suffix
}
