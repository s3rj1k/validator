package validators

import (
	"fmt"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/s3rj1k/validator"
)

// UUIDIsPresentError is a function that defines error message returned by UUIDIsPresent validator.
// nolint: gochecknoglobals
var UUIDIsPresentError = func(v *UUIDIsPresent) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must not be blank", v.Name)
}

// UUIDIsPresent is a validator object.
// Validate adds an error if the Field is an uuid default value (uuid.Nil).
type UUIDIsPresent struct {
	Name    string
	Field   uuid.UUID
	Message string
}

// Validate adds an error if the Field is an uuid default value (uuid.Nil).
func (v *UUIDIsPresent) Validate(e *validator.Errors) {
	s := v.Field.String()
	if strings.TrimSpace(s) != "" && v.Field != uuid.Nil {
		return
	}

	e.Add(v.Name, UUIDIsPresentError(v))
}
