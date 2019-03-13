package validators

import (
	"fmt"
	"os"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringFileIsExecutableDive(t *testing.T) {

	r := require.New(t)

	_ = os.Remove("/tmp/executable_file")

	fd, err := os.Create("/tmp/executable_file") // nolint: gosec
	r.Nil(err)
	err = fd.Chmod(0777)
	r.Nil(err)
	err = fd.Close()
	r.Nil(err)

	_ = os.Remove("/tmp/not_executable_file")

	fd2, err := os.Create("/tmp/not_executable_file") // nolint: gosec
	r.Nil(err)
	err = fd2.Close()
	r.Nil(err)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"/tmp/not_exists_i_hope", "/tmp/executable_file", "/tmp", "/tmp/not_executable_file"}, false, []int{3}},
		{[]string{" ", ""}, true, []int{}}, // not a file is OK
		{nil, true, []int{}},               // not a file is OK
	}

	// execs := []int{0111, 0333, 0555, 0743}
	// notExecs := []int{0000, 0700, 0770, 0222, 0444, 0666, 0743}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringFileIsExecutable{Name: "StringFileIsExecutableDive"},
			Field:     test.field,
		}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d", index)
		if !test.valid {
			r.Equalf(len(test.invalidIndexes), e.Count(), "tc %d wrong number of errors", index)

			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("StringFileIsExecutableDive[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}

	err = os.Remove("/tmp/executable_file")
	r.Nil(err)

	err = os.Remove("/tmp/not_executable_file")
	r.Nil(err)
}
