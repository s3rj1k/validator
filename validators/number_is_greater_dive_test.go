package validators

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

type numTestCase struct {
	name          string
	field         interface{}
	comparedField interface{}
	checkEqual    bool
	errNum        int
}

func Test_NumberIsGreaterDive(t *testing.T) {

	testCases := []numTestCase{
		{
			name:          "SliceOne",
			field:         []int{-1000, -1, 0, 1, 1000},
			comparedField: int(-2000),
		},
		{
			name:          "SliceTwo",
			field:         []int32{-1000, -100, 0, 1, 1000},
			comparedField: int16(-10),
			errNum:        2, // -1000 and -100 are not greater than -10
		},
		{
			name:          "SliceThree",
			field:         []int16{-200, -100, 0, 1, 20, 1000, 2000, 200, 200, 200},
			comparedField: uint8(200),
			checkEqual:    true,
			errNum:        5, /// -200, -100, 0, 1, 20 are not greater than 200
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
			Validator: &NumberIsGreater{
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
