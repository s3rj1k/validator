package validators

import (
	"fmt"
	"os"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsNotFileDive(t *testing.T) {

	r := require.New(t)

	fd, err := os.Create("/tmp/string_is_not_file_dive") // nolint: gosec
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"/tmp/not_exists_i_hope", "/tmp/string_is_not_file_dive", "/tmp"}, false, []int{1}},
		{[]string{" ", ""}, true, []int{}}, // not a file
		{nil, true, []int{}},               // not a file
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsNotFile{Name: "StringIsNotFile"},
			Field:     test.field,
		}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d", index)
		if !test.valid {
			r.Equalf(len(test.invalidIndexes), e.Count(), "tc %d wrong number of errors", index)

			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("StringIsNotFile[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}

	err = os.Remove("/tmp/string_is_not_file_dive")
	r.Nil(err)
}
