package validators

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsSymlinkDive(t *testing.T) {

	r := require.New(t)

	_ = os.Remove("/tmp/test_symlink")
	err := os.Symlink("/tmp", "/tmp/test_symlink")
	r.Nil(err)

	_ = os.Remove("/tmp/test_symlink2")
	err = os.Symlink("/tmp", "/tmp/test_symlink2")
	r.Nil(err)

	field := []string{"/tmp/test_symlink", "/tmp/test_symlink2"}

	v := StringSliceDive{
		Validator: &StringIsSymlink{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	err = os.Remove("/tmp/test_symlink")
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

	err = os.Remove("/tmp/test_symlink2")
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
