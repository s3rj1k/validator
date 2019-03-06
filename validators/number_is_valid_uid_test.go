package validators

import (
	"fmt"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_NumberIsValidUID(t *testing.T) {

	r := require.New(t)

	testCases := []numTestCase{
		{
			name:  "UID",
			field: int32(1000),
		},
		{
			name:  "UID",
			field: uint32(1005),
		},
		{
			name:   "UID",
			field:  uint32(9999999),
			errNum: 1,
		},
		{
			name:   "UID",
			field:  uint16(200),
			errNum: 1,
		},
		{
			name:   "UID",
			field:  int16(-200),
			errNum: 1,
		},
		{
			name:   "UID",
			field:  int(-1),
			errNum: 0,
		},
		{
			name:   "UID",
			field:  nil,
			errNum: 1,
		},
		{
			name:   "UID",
			field:  int64(9),
			errNum: 1,
		},
		{
			name:   "UID",
			field:  "bad type", // other than nubmer types is wrong
			errNum: 1,
		},
		{
			name:   "UID",
			field:  int16(0),
			errNum: 1,
		},
	}

	for index, tc := range testCases {

		v := NumberIsValidUID{
			Name:  tc.name,
			Field: tc.field,
		}

		e := validator.NewErrors()
		v.Validate(e)
		r.Equal(tc.errNum, e.Count(), fmt.Sprintf("tc %d element is wrong %v", index, e))
	}
}
