package buildin

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsSymlinkDive(t *testing.T) {
	r := require.New(t)

	_ = os.Remove(symlink)
	err := os.Symlink("/tmp", symlink)
	r.Nil(err)

	_ = os.Remove(anotherSymlink)
	err = os.Symlink("/tmp", anotherSymlink)
	r.Nil(err)

	field := []string{symlink, anotherSymlink}

	v := StringSliceDive{
		Validator: &StringIsSymlink{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	err = os.Remove(symlink)
	r.Nil(err)

	v = StringSliceDive{
		Validator: &StringIsSymlink{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())

	err = os.Remove(anotherSymlink)
	r.Nil(err)

	v = StringSliceDive{
		Validator: &StringIsSymlink{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(2, e.Count())
}
