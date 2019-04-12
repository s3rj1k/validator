package main

import (
	"fmt"

	v "github.com/s3rj1k/validator"
	vv "github.com/s3rj1k/validator/buildin"
)

func main() {

	// each built-in validator has standard error message
	e := v.Validate(
		&vv.StringIsAlpha{
			Name:  "Test",
			Field: "N0tAlpha",
		},
	)
	if e != nil {
		fmt.Println(e) // Output: {"Test":["'N0tAlpha' must contain only letters"]}
	}

	// value of Message field of any built-in validator will be the error message
	e = v.Validate(
		&vv.StringIsAlpha{
			Name:    "Test",
			Field:   "N0tAlpha",
			Message: "my custom message",
		},
	)
	if e != nil {
		fmt.Println(e) // Output: {"Test":["'N0tAlpha' must contain only letters"]}
	}

	// if you need to use the actual field values in error message of built-in validators
	// redefine proper Error variable
	vv.StringIsAlphaError = func(v *vv.StringIsAlpha) string {
		return fmt.Sprintf("Name %s Value %s failed validation", v.Name, v.Field)
	}

	e = v.Validate(
		&vv.StringIsAlpha{
			Name:  "Test",
			Field: "N0tAlpha",
		},
	)
	if e != nil {
		fmt.Println(e) // Output: {"Test":["Name Test Value N0tAlpha failed validation"]}
	}
}
