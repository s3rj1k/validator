package buildin

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsAbsPathDive(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field          []string
		valid          bool
		invalidIndexes []int
	}{
		{[]string{"/var/", "/usr/local/bin", regularFile, "/"}, true, []int{}},
		{[]string{"/tmp//test/test", "./test", "test", "//", "/tmp/test//"}, false, []int{0, 1, 2, 3, 4}},
		{[]string{" ", ""}, false, []int{0, 1}},
		{nil, false, []int{0}},
	}

	for index, test := range tests {
		v := &StringSliceDive{
			Validator: &StringIsAbsPath{Name: "AbsPath"},
			Field:     test.field,
		}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf(len(test.invalidIndexes), e.Count(), "tc %d wrong number of errors", index)

			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("AbsPath[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}
}
