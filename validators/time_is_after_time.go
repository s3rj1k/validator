package validators

import (
	"fmt"
	"time"

	"github.com/s3rj1k/validator"
)

// TimeIsAfterTimeError is a function that defines error message returned by TimeIsAfterTime validator.
// nolint: gochecknoglobals
var TimeIsAfterTimeError = func(v *TimeIsAfterTime) string {
	if len(v.ComparedName) == 0 {
		return fmt.Sprintf("'%s' must be after '%s'", v.Field, v.ComparedField)
	}

	return fmt.Sprintf("'%s' must be after '%s'", v.Name, v.ComparedName)
}

// TimeIsAfterTime is a validator object.
type TimeIsAfterTime struct {
	Name          string
	Field         time.Time
	ComparedName  string
	ComparedField time.Time
}

// Validate adds an error if the Field time is not after the ComparedField time.
func (v *TimeIsAfterTime) Validate(e *validator.Errors) {
	if v.Field.UnixNano() > v.ComparedField.UnixNano() {
		return
	}

	e.Add(v.Name, TimeIsAfterTimeError(v))
}
