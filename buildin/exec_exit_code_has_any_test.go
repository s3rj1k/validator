package buildin

import (
	"testing"
	"time"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_ExecExitCodeHasAny(t *testing.T) {

	r := require.New(t)

	var tests = []struct {
		command   []string
		exitCodes []int
		timeout   time.Duration
		valid     bool
	}{
		{[]string{"false"}, []int{0, 127, 1}, 0, true},
		{[]string{"true"}, []int{0, 127, 4}, 0, true},
		{[]string{"sh", "-c", "'exit'"}, []int{0, 127}, 0, true},
		{[]string{"touch", "-h"}, []int{0, 127, 1}, 0, true},

		{[]string{"false"}, []int{0, 127, 69}, 0, false},
		{[]string{"true"}, []int{69, 127, 4}, 0, false},
		{[]string{"sh", "-c", "'exit 1'"}, []int{0}, 0, false},
		{[]string{"touch", "-h"}, []int{0, 127, 69}, 0, false},
	}

	for index, test := range tests {
		v := &ExecExitCodeHasAny{Name: "Test", Command: test.command, ExitCodes: test.exitCodes}
		e := validator.NewErrors()
		v.Validate(e)

		r.Equalf(!test.valid, e.HasAny(), "tc %d expecting error=%v got=%v", index, !test.valid, e.HasAny())
		if !test.valid {
			r.Equalf([]string{ExecExitCodeHasAnyError(v)}, e.Get(v.Name), "tc %d", index)
		}
	}
}
