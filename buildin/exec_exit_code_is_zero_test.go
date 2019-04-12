package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_ExecExitCodeIsZero(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		command []string
		timeout int64
		valid   bool
	}{
		{[]string{"test", "-f", "/bin/sh"}, 0, true},
		{[]string{"test", "-d", "/root"}, 0, true},

		{[]string{"test", "-f", "/bin/non-existent-elf"}, 0, false},
		{[]string{"test", "-d", "/non-existent-dir"}, 0, false},

		{[]string{"sleep", "5"}, 1, false},
	}

	for index, test := range tests {
		v := &ExecExitCodeIsZero{Name: "Test", Command: test.command, TimeoutSeconds: test.timeout}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())
		if !test.valid {
			r.Equalf([]string{ExecExitCodeIsZeroError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
