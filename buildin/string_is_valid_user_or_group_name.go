package buildin

import (
	"fmt"

	"github.com/s3rj1k/validator"
)

// StringIsValidUserOrGroupNameError is a function that defines error message returned by StringIsValidUserOrGroupName validator.
// nolint: gochecknoglobals
var StringIsValidUserOrGroupNameError = func(v *StringIsValidUserOrGroupName) string {
	if len(v.Message) > 0 {
		return v.Message
	}

	var caseName string

	if v.CaseInsensitive {
		caseName = "name (case insensitive)"
	} else {
		caseName = "name"
	}

	return fmt.Sprintf("'%s' is not a valid user or group %s", v.Field, caseName)
}

// StringIsValidUserOrGroupName is a validator object.
// Validate adds an error if the Field is not a valid user or group name.
type StringIsValidUserOrGroupName struct {
	Name            string
	Field           string
	Message         string
	CaseInsensitive bool
}

// Validate adds an error if the Field is not a valid user or group name.
func (v *StringIsValidUserOrGroupName) Validate(e *validator.Errors) {
	if v.CaseInsensitive {
		if isValidUserOrGroupNameCaseInsensitive(v.Field) {
			return
		}
	} else {
		if isValidUserOrGroupName(v.Field) {
			return
		}
	}

	e.Add(v.Name, StringIsValidUserOrGroupNameError(v))
}

// SetField sets validator field.
func (v *StringIsValidUserOrGroupName) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsValidUserOrGroupName) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", RxSetNameIndex.ReplaceAllString(v.Name, ""), i)
}

func isValidUserOrGroupName(name string) bool {
	if len(name) < 1 || len(name) > 32 {
		return false
	}

	if !rxUserGroupName.MatchString(name) {
		return false
	}

	return true
}

func isValidUserOrGroupNameCaseInsensitive(name string) bool {
	if len(name) < 1 || len(name) > 32 {
		return false
	}

	if !rxUserGroupNameI.MatchString(name) {
		return false
	}

	return true
}
