package validators

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/s3rj1k/validator"
)

// StringIsValidShadowPasswordError is a function that defines error message returned by StringIsValidShadowPassword validator.
// nolint: gochecknoglobals
var StringIsValidShadowPasswordError = func(v *StringIsValidShadowPassword) string {
	return fmt.Sprintf("'%s' is not a valid shadow password", v.Field)
}

// StringIsValidShadowPassword is a validator object.
type StringIsValidShadowPassword struct {
	Name  string
	Field string
}

// Validate adds an error if the Field is not a valid shadow password.
func (v *StringIsValidShadowPassword) Validate(e *validator.Errors) {

	if isValidShadowPasswd(v.Field) {
		return
	}

	e.Add(v.Name, StringIsValidShadowPasswordError(v))
}

// SetField sets validator field.
func (v *StringIsValidShadowPassword) SetField(s string) {
	v.Field = s
}

// SetNameIndex sets index of slice element on Name.
func (v *StringIsValidShadowPassword) SetNameIndex(i int) {
	v.Name = fmt.Sprintf("%s[%d]", regexp.MustCompile(`\[[0-9]+\]$`).ReplaceAllString(v.Name, ""), i)
}

func isValidShadowPasswd(passwd string) bool {

	hashLength := map[string]int{
		"1": 22, // md5
		"5": 43, // sha256
		"6": 86, // sha512
	}

	// general cases
	switch passwd {
	case "", "NP", "x": // empty password is valid
		return true
	case "!", "*LK*", "*", "!!": // valid cases
		return true
	}

	splitted := strings.Split(passwd, "$")

	// simple case
	if len(splitted) == 1 {
		hash := strings.TrimPrefix(passwd, "!") // "!" at the beginning is valid

		return len(hash) == 13
	}

	// complex case
	id := splitted[1]

	hash := splitted[len(splitted)-1]
	hash = strings.TrimPrefix(hash, "!") // "!" at the beginning is valid

	if !rxPasswd.MatchString(hash) {
		return false
	}

	length, ok := hashLength[id]

	switch ok {
	case true:
		if len(hash) != length {
			return false
		}

	case false:
		if !strings.Contains(id, "2") { // other alg is blowfish, 2/2a/2b/2x/2y
			return false
		}
	}

	return true
}
