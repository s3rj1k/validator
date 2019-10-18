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

	_ = os.Remove(writableFile)

	fd, err := os.Create(writableFile)
	r.Nil(err)

	err = fd.Chmod(0777)
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	_ = os.Remove(notWritableFile)

	fd, err = os.Create(notWritableFile)
	r.Nil(err)

	err = fd.Chmod(0000)
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	var tests = []struct {
		field          []string
		invalidIndexes []int
	}{
		{[]string{notExists, notWritableFile, writableFile}, []int{0, 1}},
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

	err = os.Chmod(notWritableFile, 0777)
	r.Nil(err)

	err = os.Remove(notWritableFile)
	r.Nil(err)

	err = os.Chmod(writableFile, 0777)
	r.Nil(err)

	err = os.Remove(writableFile)
	r.Nil(err)
}
