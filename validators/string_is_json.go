package validators

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsJSONError is a function that defines error message returned by StringIsJSON validator.
// nolint: gochecknoglobals
var StringIsJSONError = func(v *StringIsJSON) string {
	return fmt.Sprintf("'%s' must be a valid JSON", v.Field)
}

// StringIsJSON is a validator object.
type StringIsJSON struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not a valid JSON.
func (v *StringIsJSON) Validate(e *validator.Errors) {

	var js json.RawMessage

	// successful unmarshalling is good
	err := json.Unmarshal([]byte(v.Field), &js)
	if err == nil {
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
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}
