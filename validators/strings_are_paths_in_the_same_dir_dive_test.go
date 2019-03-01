package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringsArePathsInTheSameDirDive(t *testing.T) {

	r := require.New(t)

	sl := []string{"Foo", "/Bar", "/tmp/Bob", "//bin//What?"}

	e := validator.NewErrorsP()
	v := StringSliceDive{
		Validator: &StringsArePathsInTheSameDir{Name: "Paths", ComparedField: "Bar"},
		Field:     sl,
	}
	v.Validate(e)
	r.Equal(3, e.Count())
}
