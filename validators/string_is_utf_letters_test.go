package validators

import (
	"testing"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_StringIsUTFLetters(t *testing.T) {

	r := require.New(t)

	v := &StringIsUTFLetters{Name: "Name", Field: "asd品ʂля"}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsUTFLetters{Name: "Name", Field: ""} // empty string is valid
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &StringIsUTFLetters{Name: "Name", Field: "123~$"} // any other characters except for UTF letters are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsUTFLettersError(v)}, e.Get("Name"))

	v = &StringIsUTFLetters{Name: "Name", Field: " ля 品ʂ "} // inner/outer whitespaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsUTFLettersError(v)}, e.Get("Name"))

	v = &StringIsUTFLetters{Name: "Name", Field: "   "} // only whitespaces are not allowed
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{StringIsUTFLettersError(v)}, e.Get("Name"))
}
