package validators

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_NumberIsNotZero(t *testing.T) {

	r := require.New(t)

	for _, n := range nonzeros2 {
		v := &NumberIsNotZero{Name: "Number", Field: n}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(0, e.Count())
	}

	for _, n := range zeros {
		v := &NumberIsNotZero{Name: "Number", Field: n}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(1, e.Count())
		r.Equal([]string{NumberIsNotZeroError(v)}, e.Get("Number"))
	}

	for _, n := range randomTypes {
		v := &NumberIsNotZero{Name: "Number", Field: n}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(1, e.Count())
		r.Equal([]string{"Number unsupported number type"}, e.Get("Number"))
	}
}
