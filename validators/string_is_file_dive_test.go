package validators

import (
	"fmt"
	"os"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsFileDive(t *testing.T) {

	r := require.New(t)

	fd, err := os.Create("/tmp/string_is_file_dive") // nolint: gosec
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"/tmp/not_exists_i_hope", "/tmp/string_is_file_dive", "/tmp"}, false, []int{0, 2}},
		{[]string{" ", ""}, false, []int{0, 1}}, // not a files
		{nil, false, []int{0}},                  // not a file
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsFile{Name: "StringIsFile"},
			Field:     test.field,
		}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())
		
		if !test.valid {
			r.Equalf(len(test.invalidIndexes), e.Count(), "tc %d wrong number of errors", index)

			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("StringIsFile[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}

	err = os.Remove("/tmp/string_is_file_dive")
	r.Nil(err)
}
