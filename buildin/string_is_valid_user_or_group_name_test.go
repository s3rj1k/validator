package buildin

import (
	"strings"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsValidUserOrGroupNameCaseSensative(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"_asdrr45u3$", true},
		{"aj8845n__--4", true},
		{"aJ8845n__--4", false},
		{"kandljv$v", false},
		{strings.Repeat("a", 33), false},
		{"a", true},
		{"asdfaeeag8gG.", false},
		{"887800924t5802", false},
		{"-aasdkllk56oh$", false},
		{" ", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsValidUserOrGroupName{Name: "Passwd", Field: test.field}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsValidUserOrGroupNameError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}

func Test_StringIsValidUserOrGroupNameCaseInsensative(t *testing.T) {
	r := require.New(t)

	var tests = []struct {
		field string
		valid bool
	}{
		{"_asdrr45u3$", true},
		{"aj8845n__--4", true},
		{"aJ8845n__--4", true},
		{"kandljv$v", false},
		{strings.Repeat("a", 33), false},
		{"a", true},
		{"asdfaeeag8gG.", false},
		{"887800924t5802", false},
		{"-aasdkllk56oh$", false},
		{" ", false},
		{"", false},
	}

	for index, test := range tests {
		v := &StringIsValidUserOrGroupName{Name: "Passwd", Field: test.field, CaseInsensitive: true}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())

		if !test.valid {
			r.Equalf([]string{StringIsValidUserOrGroupNameError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
