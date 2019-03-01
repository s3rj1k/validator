package validators

import (
	"testing"
	"time"

	"github.com/s3rj1k/validator"

	"github.com/stretchr/testify/require"
)

func Test_TimeIsPresent(t *testing.T) {

	r := require.New(t)

	v := TimeIsPresent{Name: "CreatedAt", Field: time.Now()}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = TimeIsPresent{Name: "CreatedAt", Field: time.Time{}}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"CreatedAt can not be blank"}, e.Get("CreatedAt"))

	v = TimeIsPresent{"CreatedAt", time.Time{}}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{"CreatedAt can not be blank"}, e.Get("CreatedAt"))
}
