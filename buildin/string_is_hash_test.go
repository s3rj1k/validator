package buildin

import (
	"crypto/md5" // nolint: gosec
	"crypto/sha256"
	"fmt"
	"io"
	"testing"

	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_StringIsHash(t *testing.T) {
	r := require.New(t)

	h := md5.New() // nolint: gosec

	_, err := io.WriteString(h, "Hello World!")
	if err != nil {
		t.Fatalf("unexpected error %s", err)
	}

	v := &StringIsHash{Name: "Name", Algorithm: "md5", Field: fmt.Sprintf("%x", h.Sum(nil))} // md5
	e := validator.NewErrors()

	v.Validate(e)
	r.Equal(0, e.Count())

	h = sha256.New()

	_, err = io.WriteString(h, "Hello World!")
	if err != nil {
		t.Fatalf("unexpected error %s", err)
	}

	v = &StringIsHash{Name: "Name", Algorithm: "sha256", Field: fmt.Sprintf("%x", h.Sum(nil))} // sha256
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsHash{Name: "Name", Algorithm: "", Field: fmt.Sprintf("%x", h.Sum(nil))} // empty algorithm is invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsHashError(v)}, e.Get("Name"))

	v = &StringIsHash{Name: "Name", Algorithm: "unknown", Field: fmt.Sprintf("%x", h.Sum(nil))} // unknown algorithm is invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsHashError(v)}, e.Get("Name"))

	v = &StringIsHash{Name: "Name", Algorithm: "md5", Field: "whatisthis"} // random string is invalid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsHashError(v)}, e.Get("Name"))
}
