package buildin

import (
	"fmt"
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsPathAndIsWritableDive(t *testing.T) {
	r := require.New(t)

	_ = os.Remove("/tmp/string_writable_file")

	fd, err := os.Create("/tmp/string_writable_file")
	r.Nil(err)

	err = fd.Chmod(0777)
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	_ = os.Remove("/tmp/string_not_writable_file")

	fd, err = os.Create("/tmp/string_not_writable_file")
	r.Nil(err)

	err = fd.Chmod(0000)
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	var tests = []struct {
		field          []string
		invalidIndexes []int
	}{
		{[]string{"/tmp/not_exists_i_hope", "/tmp/string_not_writable_file", "/tmp/string_writable_file"}, []int{0, 1}},
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsPathAndIsWritable{Name: "StringIsPathAndIsWritableDive"},
			Field:     test.field,
		}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(len(test.invalidIndexes), e.Count(), "tc %d wrong number of errors", index)

		errnames := []string{}
		for _, i := range test.invalidIndexes {
			errnames = append(errnames, fmt.Sprintf("StringIsPathAndIsWritableDive[%d]", i))
		}

		for _, en := range errnames {
			r.Containsf(e.JSON(), en, "tc %d", index)
		}
	}

	err = os.Chmod("/tmp/string_not_writable_file", 0777)
	r.Nil(err)

	err = os.Remove("/tmp/string_not_writable_file")
	r.Nil(err)

	err = os.Chmod("/tmp/string_writable_file", 0777)
	r.Nil(err)

	err = os.Remove("/tmp/string_writable_file")
	r.Nil(err)
}
