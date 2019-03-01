package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsPrintableASCIIDive(t *testing.T) {

	r := require.New(t)

	field := []string{"asd456", "0nlY", " ASC11", "!$#%()-=<>etc...,@", " ", ""}

	v := StringSliceDive{
		Validator: &StringIsPrintableASCII{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"тут", "ошибочка", " s0m3", "oth &r", "sym bols", " ", "", string(rune(10))}

	v = StringSliceDive{
		Validator: &StringIsPrintableASCII{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(3, e.Count())
}
