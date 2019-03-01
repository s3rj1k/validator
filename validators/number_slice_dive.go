package validators

import (
	"github.com/s3rj1k/validator"
)

// NumberValidator is an interface for string validator objects.
type NumberValidator interface {
	Validate(*validator.Errors)
	SetField(interface{})
	SetNameIndex(int)
}

// NumberSliceDive is a validator object
type NumberSliceDive struct {
	Validator NumberValidator
	Field     []interface{}
}

// Validate applies Validator to each value in the Field.
func (v *NumberSliceDive) Validate(e *validator.Errors) {
	for i, f := range v.Field {
		v.Validator.SetField(f)
		v.Validator.SetNameIndex(i)
		v.Validator.Validate(e)
	}
}
