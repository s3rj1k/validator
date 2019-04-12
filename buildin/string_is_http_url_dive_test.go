package buildin

import (
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsHTTPURLDive(t *testing.T) {

	r := require.New(t)

	field := []string{"", "http://", "https://", "http", "google.com",
		"http://www.google.com", "http://google.com", "https://www.google.cOM",
		"ht123tps://www.google.cOM",
		"https://golang.Org",
		"https://invalid#$%#$@.Org"}

	v := StringSliceDive{
		Validator: &StringIsHTTPURL{
			Name: "MySlice",
		},
		Field: field,
	}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(7, e.Count())
}
