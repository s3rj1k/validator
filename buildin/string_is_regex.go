package buildin

import (
	"fmt"
	"regexp"

	"github.com/s3rj1k/validator"
)

// StringIsRegexError is a function that defines error message returned by StringIsRegex validator.
// nolint: gochecknoglobals
var StringIsRegexError = func(v *StringIsRegex) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	if v.POSIX {
		return fmt.Sprintf("'%s' is not valid POSIX regular expression", v.Field)
	}

	return fmt.Sprintf("'%s' is not valid regular expression", v.Field)
}

// StringIsRegex is a validator object.
// Validate adds an error if the Field is not valid regular expression.
type StringIsRegex struct {
	Name    string
	Field   string
	Message string
	POSIX   bool
}

// Validate adds an error if the Field is not valid regular expression.
func (v *StringIsRegex) Validate(e *validator.Errors) {
	if v.POSIX {
		if _, err := regexp.CompilePOSIX(v.Field); err != nil {
			e.Add(v.Name, StringIsRegexError(v))
		}
	} else {
		if _, err := regexp.Compile(v.Field); err != nil {
			e.Add(v.Name, StringIsRegexError(v))
		}
	}
}

// SetField sets validator field.
func (v *StringIsRegex) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsRegex) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
