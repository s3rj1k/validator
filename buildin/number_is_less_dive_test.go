package buildin

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_NumberIsLessDive(t *testing.T) {
	testCases := []numTestCase{
		{
			name:          "0",
			field:         []int{-10000, -9000, 0, 1, 1000},
			comparedField: int(1001),
		},
		{
			name:           "1",
			field:          []int32{-1000, -100, 0, 1, 1000},
			comparedField:  int16(-99),
			invalidIndexes: []int{2, 3, 4}, // 0, 1, 1000 are not less than -99
		},
		{
			name:           "2",
			field:          []int16{-200, -100, 0, 1, 20, 1000, 2000, 200, 200, 200, 200},
			comparedField:  uint8(200),
			checkEqual:     true,
			invalidIndexes: []int{5, 6}, /// 1000, 2000 are not less than 200
		},
		{
			name:           "3",
			field:          nil, // nil field = []int8{0}
			comparedField:  int16(0),
			invalidIndexes: []int{0},
		},
		{
			name:           "4",
			field:          []int32{-1000, -1, 0, 1, 1000},
			comparedField:  nil, // nil field = []int8{0}
			invalidIndexes: []int{2, 3, 4},
		},
		{
			name:           "5",
			field:          "bad type", // other than nubmer types is wrong
			comparedField:  int16(0),
			invalidIndexes: []int{0},
		},
		{
			name:           "6",
			field:          []int32{-1000, -1, 0, 1, 1000},
			comparedField:  "bad type", // other than nubmer types is wrong. will add error for each value in field
			invalidIndexes: []int{0, 1, 2, 3, 4},
		},
	}

	r := require.New(t)

	for index, tc := range testCases {
		v := NumberSliceDive{
			Validator: &NumberIsLess{
				Name:          tc.name,
				ComparedField: tc.comparedField,
				CheckEqual:    tc.checkEqual,
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
