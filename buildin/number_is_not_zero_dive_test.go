package buildin

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_NumberIsNotZeroDive(t *testing.T) {
	testCases := []numTestCase{
		{
			name:  "0",
			field: []int{-1000, -1, 1, 1000},
		},
		{
			name:           "1",
			field:          []int32{-1000, -100, 0, 1, 1000},
			invalidIndexes: []int{2},
		},
		{
			name:           "2",
			field:          []int16{-200, -100, 0, 0, 20, 1000, 2000},
			invalidIndexes: []int{2, 3},
		},
		{
			name:           "3",
			field:          nil, // nil field = []int8{0}
			invalidIndexes: []int{0},
		},
		{
			name:           "4",
			field:          "bad type", // other than nubmer types is wrong
			invalidIndexes: []int{0},
		},
	}

	r := require.New(t)

	for index, tc := range testCases {
		v := NumberSliceDive{
			Validator: &NumberIsNotZero{
				Name: tc.name,
			},
			Field: tc.field,
		}
		e := validator.NewErrors()

		v.Validate(e)

		// check number of errors
		r.Equal(len(tc.invalidIndexes), e.Count(), fmt.Sprintf("tc '%d' expecting '%d' errors got '%d'", index, len(tc.invalidIndexes), e.Count()))

		if len(tc.invalidIndexes) > 0 {
			// case when Field is of a wrong type
			if len(e.Get(tc.name)) == 1 && e.Get(tc.name)[0] == ErrBadNumType.Error() {
				break
			}

			// check that final errors contain all indexes provided in invalid indexes
			errnames := []string{}
			for _, i := range tc.invalidIndexes {
				errnames = append(errnames, fmt.Sprintf("%s[%d]", tc.name, i))
			}

			for _, en := range errnames {
				r.Containsf(e.JSON(), en, "tc %d", index)
			}
		}
	}
}
