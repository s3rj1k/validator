package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringsArePathsNotInTheSameDir(t *testing.T) {
	r := require.New(t)

	var cases = []struct {
		str1     string
		str2     string
		expected bool
	}{
		{"test", "test", false},
		{"ololo/test_fail", "zzzz/test_true", true},
		{"/sdgsdg/sdgsdg/test with space", "/sdgsdg/test with space ", true},
		{"/sss/test with space second", "/sss/test with space second       ", false},
	}

	for _, testCase := range cases {
		v := StringsArePathsNotInTheSameDir{Name: "paths", Field: testCase.str1, ComparedField: testCase.str2}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Path1: %s, Path2: %s", testCase.str1, testCase.str2)
	}

	for _, testCase := range cases {
		v := StringsArePathsNotInTheSameDir{Name: "Path1", Field: testCase.str1, ComparedField: testCase.str2, ComparedName: "Path2"}
		e := validator.NewErrors()

		v.Validate(e)
		r.Equal(testCase.expected, !e.HasAny(), "Str1: %s, Str2: %s", testCase.str1, testCase.str2)

		if !testCase.expected {
			r.Contains(e.Get("Path1"), "'Path1' path is in the same dir with 'Path2'")
		}
	}
}
