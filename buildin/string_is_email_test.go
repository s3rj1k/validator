package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsEmail(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		email string
		valid bool
	}{
		{"", false},
		{"foo@bar.com", true},
		{"x@x.x", true},
		{"foo@bar.com.au", true},
		{"foo+bar@bar.com", true},
		{"foo@bar.coffee", true},
		{"foo@bar.中文网", true},
		{"invalidemail@", false},
		{"invalid.com", false},
		{"@invalid.com", false},
		{"test|123@m端ller.com", true},
		{"hans@m端ller.com", true},
		{"hans.m端ller@test.com", true},
		{"NathAn.daVIeS@DomaIn.cOM", true},
		{"NATHAN.DAVIES@DOMAIN.CO.UK", true},
	}

	for _, test := range tests {
		v := &StringIsEmail{Name: "email", Field: test.email}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equal(test.valid, !e.HasAny())

		if !test.valid {
			r.Equal([]string{StringIsEmailError(v)}, e.Get("email"))
		}
	}

	v := &StringIsEmail{Name: "email", Field: ""}
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(e.Count(), 1)
	r.Equal([]string{StringIsEmailError(v)}, e.Get("email"))
}
