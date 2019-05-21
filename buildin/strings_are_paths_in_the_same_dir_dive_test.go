package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringsArePathsInTheSameDirDive(t *testing.T) {
	r := require.New(t)

	sl := []string{"Foo", "/Bar", "/tmp/Bob", "//bin//What?"}

	v := StringSliceDive{
		Validator: &StringsArePathsInTheSameDir{Name: "Paths", ComparedField: "Bar"},
		Field:     sl,
	}
	e := validator.NewErrorsP()
	v.Validate(e)
	r.Equal(3, e.Count())
}
