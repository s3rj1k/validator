package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_FuncValidator(t *testing.T) {

	r := require.New(t)

	fv := &FuncValidator{
		Name:  "CustomFunc",
		Field: "Field",
		Fn: func() bool {
			return false
		},
	}

	e := validator.NewErrors()
	fv.Validate(e)
	r.Equal([]string{FuncValidatorError(fv)}, e.Get("CustomFunc"))
}
