package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringExclusionError is a function that defines error message returned by StringExclusion validator.
// nolint: gochecknoglobals
var StringExclusionError = func(v *StringExclusion) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("%s is in the blacklist %v", v.Name, v.Blacklist)
}

// StringExclusion is a validator object.
// Validate adds an error if the Field is one of the values from the Blacklist.
type StringExclusion struct {
	Name      string
	Field     string
	Blacklist []string
	Message   string
}

// Validate adds an error if the Field is one of the values from the Blacklist.
func (v *StringExclusion) Validate(e *validator.Errors) {
	var found = false

	for _, l := range v.Blacklist {
		if l == v.Field {
			found = true

			break
		}
	}

	if found {
		e.Add(v.Name, StringExclusionError(v))
	}
}

// SetField sets validator field.
func (v *StringExclusion) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringExclusion) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
