package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsEmailLikeDive(t *testing.T) {

	r := require.New(t)

	field := []string{"", " foo@bar.com   ", "x@x.x", "foo  @bar.com.au", "foo+bar@bar.com", "foo@bar.coffee",
		"foo@bar.中文网", "invalidemail@", "invalid.com", "@", "@invalid.com", "test|123@m端ller.com", "hans@m端ller.com",
		"hans.m端ller@test.com", "NathAn.daVIeS@DomaIn.cOM", "NATHAN.DAVIES@DOMAIN.CO.UK"}

	v := StringSliceDive{
		Validator: &StringIsEmailLike{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(5, e.Count())
}
