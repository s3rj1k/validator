package validators

import (
	"testing"
	"time"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_ExecExitCodeIsNotZero(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		command []string
		timeout time.Duration
		valid   bool
	}{
		{[]string{"test", "-f", "/bin/sh"}, 0, false},
		{[]string{"test", "-d", "/root"}, 0, false},

		{[]string{"test", "-f", "/bin/non-existent-elf"}, 0, true},
		{[]string{"test", "-d", "/non-existent-dir"}, 0, true},

		// {[]string{"sleep", "5"}, 1 * time.Millisecond, true},
	}

	for index, test := range tests {
		v := &ExecExitCodeIsNotZero{Name: "Test", Command: test.command}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())
		if !test.valid {
			r.Equalf([]string{ExecExitCodeIsNotZeroError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
