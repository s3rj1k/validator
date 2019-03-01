package validators

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_NumberIsGreater(t *testing.T) {

	r := require.New(t)

	for i := range nonzeros2 {
		v := &NumberIsGreater{Name: "Number", Field: nonzeros10[i], ComparedField: nonzeros2[i]}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(0, e.Count())
	}

	for i := range nonzeros2 {
		v := &NumberIsGreater{Name: "Number", Field: nonzeros2[i], ComparedField: nonzeros10[i]}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(1, e.Count())
		r.Equal([]string{NumberIsGreaterError(v)}, e.Get("Number"))
	}

	for _, n := range randomTypes {
		v := &NumberIsGreater{Name: "Number", Field: n}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(1, e.Count())
		r.Equal([]string{"Number nil fields are forbidden"}, e.Get("Number"))
	}

	for _, n := range randomTypes {
		v := &NumberIsGreater{Name: "Number", Field: n, ComparedField: n}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(1, e.Count())
		r.Equal([]string{"Number types cannot be compared"}, e.Get("Number"))
	}
}
