package buildin

import (
	"encoding/json"
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsJSONError is a function that defines error message returned by StringIsJSON validator.
// nolint: gochecknoglobals
var StringIsJSONError = func(v *StringIsJSON) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' must be a valid JSON", v.Field)
}

// StringIsJSON is a validator object.
// Validate adds an error if the Field is not a valid JSON.
type StringIsJSON struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is not a valid JSON.
func (v *StringIsJSON) Validate(e *validator.Errors) {
	var js json.RawMessage

	// successful unmarshalling is good
	if err := json.Unmarshal([]byte(v.Field), &js); err == nil {
		return
	}

	e.Add(v.Name, StringIsJSONError(v))
}

// SetField sets validator field.
func (v *StringIsJSON) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsJSON) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
