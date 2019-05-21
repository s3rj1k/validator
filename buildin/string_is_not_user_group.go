package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsNotUserGroupError is a function that defines error message returned by StringIsNotUserGroup validator.
// nolint: gochecknoglobals
var StringIsNotUserGroupError = func(v *StringIsNotUserGroup) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	return fmt.Sprintf("'%s' is a user group", v.Field)
}

// StringIsNotUserGroup is a validator object.
// Validate adds an error if the Field is a user group.
type StringIsNotUserGroup struct {
	Name    string
	Field   string
	Message string
}

// Validate adds an error if the Field is a user group.
func (v *StringIsNotUserGroup) Validate(e *validator.Errors) {
	if !IsGroupIsUserGroupOrWhitelisted(v.Field) {
		return
	}

	e.Add(v.Name, StringIsNotUserGroupError(v))
}

// SetField sets validator field.
func (v *StringIsNotUserGroup) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsNotUserGroup) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}
