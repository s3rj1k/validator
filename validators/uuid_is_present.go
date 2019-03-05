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
	return fmt.Sprintf("%s can not be blank", v.Name)
}

// UUIDIsPresent is a validator object
type UUIDIsPresent struct {
	Name  string
	Field uuid.UUID
}

// Validate adds an error if the Field is an uuid default value (uuid.Nil).
func (v *UUIDIsPresent) Validate(e *validator.Errors) {
	s := v.Field.String()
	if strings.TrimSpace(s) != "" && v.Field != uuid.Nil {
		return
	}

	e.Add(v.Name, UUIDIsPresentError(v))
}
