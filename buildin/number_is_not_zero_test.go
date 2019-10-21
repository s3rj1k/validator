package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_NumberIsNotZero(t *testing.T) {
	r := require.New(t)

	var zeros = []interface{}{int(0), int8(0), int16(0), int32(0), int64(0),
		uint(0), uint8(0), uint16(0), uint32(0), uint64(0),
		rune(0), byte(0)}

	var nonzeros = []interface{}{
		int(2), int8(2), int16(2), int32(2), int64(2),
		uint(2), uint8(2), uint16(2), uint32(2), uint64(2),
		rune(2), byte(2)}

	for _, n := range nonzeros {
		v := &NumberIsNotZero{Name: "Number", Field: n}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equal(0, e.Count())
	}

	for _, n := range zeros {
		v := &NumberIsNotZero{Name: "Number", Field: n}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equal(1, e.Count())
		r.Equal([]string{NumberIsNotZeroError(v)}, e.Get("Number"))
	}
}
