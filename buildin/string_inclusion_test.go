package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringInclusion(t *testing.T) {

	r := require.New(t)

	l := []string{"Mark", "Bates"}

	v := &StringInclusion{Name: "Name", Field: "Mark", Whitelist: l}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringInclusion{Name: "Name", Field: "Foo", Whitelist: l}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringInclusionError(v)}, e.Get("Name"))
}
