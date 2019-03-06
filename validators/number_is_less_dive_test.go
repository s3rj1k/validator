package validators

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_NumberIsLessDive(t *testing.T) {

	testCases := []numTestCase{
		{
			name:          "SliceOne",
			field:         []int{-10000, -9000, 0, 1, 1000},
			comparedField: int(1001),
		},
		{
			name:          "SliceTwo",
			field:         []int32{-1000, -100, 0, 1, 1000},
			comparedField: int16(-99),
			errNum:        3, // 0, 1, 1000 are not less than -99
		},
		{
			name:          "SliceThree",
			field:         []int16{-200, -100, 0, 1, 20, 1000, 2000, 200, 200, 200, 200},
			comparedField: uint8(200),
			checkEqual:    true,
			errNum:        2, /// 1000, 2000 are not less than 200
		},
		{
			name:          "SliceFour",
			field:         nil, // nil field is wrong
			comparedField: int16(0),
			errNum:        1,
		},
		{
			name:          "SliceFive",
			field:         []int32{-1000, -1, 0, 1, 1000},
			comparedField: nil, // nil comparedField is wrong. will add error for each value in field
			errNum:        5,
		},
		{
			name:          "SliceFour",
			field:         "bad type", // other than nubmer types is wrong
			comparedField: int16(0),
			errNum:        1,
		},
		{
			name:          "SliceFive",
			field:         []int32{-1000, -1, 0, 1, 1000},
			comparedField: "bad type", // other than nubmer types is wrong. will add error for each value in field
			errNum:        5,
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
		r.Equal(tc.errNum, e.Count(), fmt.Sprintf("tc %d number of errors is wrong %v", index, e))
	}
}
