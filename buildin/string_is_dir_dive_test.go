package buildin

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsDirDive(t *testing.T) {
	r := require.New(t)

	err := os.MkdirAll(dir, 0666)
	r.Nil(err)

	err = os.MkdirAll(anotherDir, 0666)
	r.Nil(err)

	field := []string{dir, anotherDir}

	v := StringSliceDive{
		Validator: &StringIsDir{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	err = os.Remove(dir)
	r.Nil(err)

	v = StringSliceDive{
		Validator: &StringIsDir{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())

	err = os.Remove(anotherDir)
	r.Nil(err)

	v = StringSliceDive{
		Validator: &StringIsDir{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(2, e.Count())
}
