package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsRegex(t *testing.T) {
	r := require.New(t)

	v := &StringIsRegex{Name: "Regexp", Field: "^([0-9]{3}-[0-9]{3}-[0-9]{4})$"}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsRegex{Name: "Regexp", Field: "^(0-9]{3}-0-9]3}-[0-9{4})$"}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsRegexError(v)}, e.Get("Regexp"))
}
