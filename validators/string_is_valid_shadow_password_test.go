package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsValidShadowPassword(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"$1$X7icamQl$jSeIrOzj3TQqwTn5kwF/50", true},
		{"$1$X7icamQl$!jSeIrOzj3TQqwTn5kwF/50", true},
		{"$1$X7icamQl$!!jSeIrOzj3TQqwTn5kwF/50", false},
		{"Ep6mckrOLChF.", true}, // https://www.tldp.org/LDP/lame/LAME/linux-admin-made-easy/shadow-file-formats.html
		{"!Ep6mckrOLChF.", true},
		{"!!Ep6mckrOLChF.", false},
		{"$6$Ke02nYgo.9v0SF4p$!hjztYvo/M4buqO4oBX8KZTftjCn6fE4cV5o/I95QPekeQpITwFTRbDUBYBLIUx2mhorQoj9bLN8v.w6btE9xy1", true},
		{"$6$Ke02nYgo.9v0SF4p$!!hjztYvo/M4buqO4oBX8KZTftjCn6fE4cV5o/I95QPekeQpITwFTRbDUBYBLIUx2mhorQoj9bLN8v.w6btE9xy1", false},
		{"", true},
		{"!", true},
		{"*", true},
	}

	for index, test := range tests {
		v := &StringIsValidShadowPassword{Name: "Passwd", Field: test.field}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())
		
		if !test.valid {
			r.Equalf([]string{StringIsValidShadowPasswordError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
