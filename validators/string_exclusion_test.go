package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringExclusion(t *testing.T) {

	r := require.New(t)

	l := []string{"Mark", "Bates"}

	v := StringExclusion{Name: "Name", Field: "Woody", Blacklist: l}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = StringExclusion{Name: "Name", Field: "Mark", Blacklist: l}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name is in the blacklist [Mark, Bates]"}, e.Get("Name"))

	v = StringExclusion{"Name", "Bates", l}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"Name is in the blacklist [Mark, Bates]"}, e.Get("Name"))
}
