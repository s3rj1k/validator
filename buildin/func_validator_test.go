package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_FuncValidator(t *testing.T) {
	r := require.New(t)

	f := func(i interface{}) bool {
		switch v := i.(type) {
		case bool:
			return v
		default:
			return false
		}
	}

	fv1 := &FuncValidator{
		Name:  "CustomFunc1",
		Field: true,
		Fn:    f,
	}

	e := validator.NewErrors()
	fv1.Validate(e)
	r.Equal(0, e.Count())

	fv2 := &FuncValidator{
		Name:  "CustomFunc2",
		Field: false,
		Fn:    f,
	}

	e = validator.NewErrors()
	fv2.Validate(e)
	r.Equal(1, e.Count())

	fv3 := &FuncValidator{
		Name:  "CustomFunc3",
		Field: "true",
		Fn:    f,
	}

	e = validator.NewErrors()
	fv3.Validate(e)
	r.Equal(1, e.Count())
}
