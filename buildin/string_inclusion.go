package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringInclusionError is a function that defines error message returned by StringInclusion validator.
// nolint: gochecknoglobals
var StringInclusionError = func(v *StringInclusion) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is not in the whitelist %v", v.Field, v.Whitelist)
}

// StringInclusion is a validator object.
// Validate adds an error if the Field is NOT one of the values from the Whitelist.
type StringInclusion struct {
	Name      string
	Field     string
	Whitelist []string
	Message   string
}

// Validate adds an error if the Field is NOT one of the values from the Whitelist.
func (v *StringInclusion) Validate(e *validator.Errors) {
	var found = false

	for _, l := range v.Whitelist {
		if l == v.Field {
			found = true
			break
		}
	}

	if !found {
		e.Add(v.Name, StringInclusionError(v))
	}
}

// SetField sets validator field.
func (v *StringInclusion) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringInclusion) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
