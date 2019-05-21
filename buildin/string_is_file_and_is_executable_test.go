package buildin

import (
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

// nolint: gosec
func Test_StringIsFileAndIsExecutable(t *testing.T) {
	r := require.New(t)

	_ = os.Remove("/tmp/string_executable_file")

	fd, err := os.Create("/tmp/string_executable_file")
	r.Nil(err)

	err = fd.Chmod(0111) // execute
	r.Nil(err)

	v := &StringIsFileAndIsExecutable{Name: "Name", Field: "/tmp/string_executable_file"} // executable is OK
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	err = fd.Chmod(0666) // read & write
	r.Nil(err)

	v = &StringIsFileAndIsExecutable{Name: "Name", Field: "/tmp/string_executable_file"} // not executable
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsFileAndIsExecutableError(v)}, e.Get("Name"))

	err = fd.Chmod(0333) // write & execute
	r.Nil(err)

	v = &StringIsFileAndIsExecutable{Name: "Name", Field: "/tmp/string_executable_file"} // not executable
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	err = fd.Chmod(0444) // read
	r.Nil(err)

	v = &StringIsFileAndIsExecutable{Name: "Name", Field: "/tmp/string_executable_file"} // not executable
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsFileAndIsExecutableError(v)}, e.Get("Name"))

	err = fd.Chmod(0555) // read & execute
	r.Nil(err)

	v = &StringIsFileAndIsExecutable{Name: "Name", Field: "/tmp/string_executable_file"} // not executable
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsFileAndIsExecutable{Name: "Name", Field: "/tmp"} // folder is not a file, no error
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsFileAndIsExecutable{Name: "Name", Field: "/tmp/not_exist_at_all"} // not existing is OK
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	err = fd.Close()
	r.Nil(err)
	err = os.Remove("/tmp/string_executable_file")
	r.Nil(err)
}
