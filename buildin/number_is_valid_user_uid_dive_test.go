package buildin

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_NumberIsValidUserUIDDive(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field          interface{}
		valid          bool
		invalidIndexes []int
	}{
		{[]int{1000, 01000, 0x1000}, false, []int{1}}, // 01000 =  512; 0x1000 = 4096
		{[]uintptr{9, 99, 999, 0x999}, false, []int{0, 1, 2}},

		{[]int32{-200, -1, 0, 1, 200}, false, []int{0, 2, 3, 4}}, // -1 is valid
		{[]uint64{59999, 60000, 60001}, false, []int{2}},
	}

	for index, test := range tests {
		v := NumberSliceDive{
			Validator: &NumberIsValidUserUID{Name: "UserUID"},
			Field:     test.field,
		}

		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			errnames := []string{}
			for _, i := range test.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("UserUID[%d]", i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}
}
