package validators

import (
	"github.com/s3rj1k/validator"
)

// StringValidator is an interface for string validator objects.
type StringValidator interface {
	Validate(*validator.Errors)
	SetField(s string)
	SetNameIndex(i int)
}

// StringSliceDive is a validator object
type StringSliceDive struct {
	Validator StringValidator
	Field     []string
}

// Validate applies Validator to each value in the Field.
func (v *StringSliceDive) Validate(e *validator.Errors) {

	slice := v.Field

	if slice == nil {
		slice = []string{""}
	}

	for i, f := range slice {
		v.Validator.SetField(f)
		v.Validator.SetNameIndex(i)
		v.Validator.Validate(e)
	}
}
