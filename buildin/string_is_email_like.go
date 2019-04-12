package buildin

import (
	"fmt"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsEmailLikeError is a function that defines error message returned by StringIsEmailLike validator.
// nolint: gochecknoglobals
var StringIsEmailLikeError = func(v *StringIsEmailLike) string {

	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("%s does not match an email-like format", v.Field)
}

// StringIsEmailLike is a validator object.
// Validate adds an error if the Field does not correspond to "username@domain" structure.
// It also checks that domain has a domain zone (but does not check if the zone is valid).
// Also allows inner and outer whitespaces.
type StringIsEmailLike struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field does not correspond to "username@domain" structure.
// It also checks that domain has a domain zone (but does not check if the zone is valid).
// Also allows inner and outer whitespaces.
func (v *StringIsEmailLike) Validate(e *validator.Errors) {

	var validStructure = false
	var domainZonePresent = false

	parts := strings.Split(v.Field, "@")

	if len(parts) == 2 && len(parts[0]) > 0 && len(parts[1]) > 0 {
		validStructure = true
	}

	if len(parts) == 2 {
		domain := parts[1]
		// Check that domain is valid
		if len(strings.Split(domain, ".")) >= 2 {
			domainZonePresent = true
		}
	}

	if !validStructure || !domainZonePresent {
		e.Add(v.Name, StringIsEmailLikeError(v))
		return
	}
}

// SetField sets validator field.
func (v *StringIsEmailLike) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsEmailLike) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
