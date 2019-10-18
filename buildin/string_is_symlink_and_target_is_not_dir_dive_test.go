package buildin

import (
	"fmt"
	"os"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsSymlinkAndTargetIsNotDirDive(t *testing.T) {
	r := require.New(t)

	fd, err := os.Create(notSymlink) // nolint: gosec
	r.Nil(err)

	err = fd.Close()
	r.Nil(err)

	_ = os.Remove(symlinkTargetIsFile)
	err = os.Symlink(notSymlink, symlinkTargetIsFile) // symlink to file
	r.Nil(err)

	_ = os.Remove(symlinkTargetIsDir)
	err = os.Symlink("/tmp", symlinkTargetIsDir) // symlink to folder
	r.Nil(err)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{notSymlink, symlinkTargetIsFile, symlinkTargetIsDir, notExists}, false, []int{2}},
		{[]string{" ", ""}, true, []int{}}, // empty and nil are not errors of this validator
		{nil, true, []int{}},
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsSymlinkAndTargetIsNotDir{Name: "SymlinkTargetIsDir"},
			Field:     test.field,
		}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf(len(test.invalidIndexes), e.Count(), "tc %d wrong number of errors", index)

			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("SymlinkTargetIsDir[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}

	err = os.Remove(notSymlink)
	r.Nil(err)

	err = os.Remove(symlinkTargetIsFile)
	r.Nil(err)

	err = os.Remove(symlinkTargetIsDir)
	r.Nil(err)
}
