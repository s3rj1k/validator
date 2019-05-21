package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsIPHasPTRDive(t *testing.T) {
	r := require.New(t)

	field := []string{"31.13.81.36", "199.59.149.230", "207.97.227.239"}

	v := StringSliceDive{
		Validator: &StringIsIPHasPTR{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	field = []string{"0.0.0.0", "192.168.255.255", "224.10.10.255", "fe8x::ffff", "10.10.10.255", " ", ""}

	v = StringSliceDive{
		Validator: &StringIsIPHasPTR{
			Name: "MySlice",
		},
		Field: field,
	}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(7, e.Count())
}
