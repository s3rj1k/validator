package validators

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_NumberIsNotZeroDive(t *testing.T) {

	testCases := []numTestCase{
		{
			name:  "SliceOne",
			field: []int{-1000, -1, 1, 1000},
		},
		{
			name:   "SliceTwo",
			field:  []int32{-1000, -100, 0, 1, 1000},
			errNum: 1, // -1000 and -100 are not greater than -10
		},
		{
			name:   "SliceThree",
			field:  []int16{-200, -100, 0, 0, 20, 1000, 2000},
			errNum: 2, /// -200, -100, 0, 1, 20 are not greater than 200
		},
		{
			name:   "SliceFour",
			field:  nil, // nil field is wrong
			errNum: 1,
		},
		{
			name:   "SliceFour",
			field:  "bad type", // other than nubmer types is wrong
			errNum: 1,
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
		r.Equal(tc.errNum, e.Count(), fmt.Sprintf("tc %d number of errors is wrong %v", index, e))
	}
}
