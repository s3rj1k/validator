package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringsArePathsNotInTheSameDirDive(t *testing.T) {

	r := require.New(t)

	sl := []string{"Foo", "/Bar", "/tmp/Bob", "//bin//What?"}

	e := validator.NewErrorsP()
	v := StringSliceDive{
		Validator: &StringsArePathsNotInTheSameDir{Name: "Paths", ComparedField: "Bar"},
		Field:     sl,
	}
	v.Validate(e)
	r.Equal(1, e.Count())
}
