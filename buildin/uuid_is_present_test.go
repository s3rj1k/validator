package buildin

import (
	"testing"

	uuid "github.com/gofrs/uuid"
	"github.com/s3rj1k/validator"
	"github.com/stretchr/testify/require"
)

func Test_UUIDIsPresent(t *testing.T) {
	r := require.New(t)

	id, err := uuid.NewV4()
	r.NoError(err)

	v := &UUIDIsPresent{Name: "Name", Field: id}
	e := validator.NewErrors()
	v.Validate(e)
	r.Equal(0, e.Count())

	v = &UUIDIsPresent{Name: "Name", Field: uuid.UUID{}}
	e = validator.NewErrors()
	v.Validate(e)
	r.Equal(1, e.Count())
	r.Equal([]string{UUIDIsPresentError(v)}, e.Get("Name"))
}
