package validators

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_NumberIsLess(t *testing.T) {

	testCases := []numTestCase{
		{
			name:          "SliceOne",
			field:         int32(100),
			comparedField: int64(1000),
		},
		{
			name:          "SliceTwo",
			field:         int16(-1000),
			comparedField: uint32(1),
		},
		{
			name:          "SliceThree",
			field:         uint16(200),
			comparedField: int32(200),
			errNum:        1, /// equal is not less
		},
		{
			name:          "SliceThree",
			field:         uint16(200),
			comparedField: int32(200),
			checkEqual:    true,
			errNum:        0,
		},
		{
			name:          "SliceFour",
			field:         nil, // nil field is wrong
			comparedField: int16(0),
			errNum:        1,
		},
		{
			name:          "SliceFive",
			field:         int64(9),
			comparedField: nil, // nil comparedField is wrong
			errNum:        1,
		},
		{
			name:          "SliceFour",
			field:         "bad type", // other than nubmer types is wrong
			comparedField: int16(0),
			errNum:        1,
		},
		{
			name:          "SliceFive",
			field:         int16(0),
			comparedField: "bad type", // other than nubmer types is wrong. will add error for each value in field
			errNum:        1,
		},
	}

	r := require.New(t)

	for index, tc := range testCases {

		v := NumberIsLess{
			Name:          tc.name,
			ComparedField: tc.comparedField,
			Field:         tc.field,
			CheckEqual:    tc.checkEqual,
		}
		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(tc.errNum, e.Count(), fmt.Sprintf("tc %d number of errors is wrong %v", index, e))
	}
}
