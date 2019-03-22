package main

import (
	"fmt"

	v "github.com/s3rj1k/validator"
	vv "github.com/s3rj1k/validator/validators"
)

func main() {

	e := v.Validate(
		&vv.StringIsAlpha{
			Name:  "Test",
			Field: "N0tAlpha",
		},
	)

	fmt.Println("Before change: ", e.JSON())

	vv.StringIsAlphaError = func(v *vv.StringIsAlpha) string {
		return fmt.Sprintf("New error message for '%s'='%s'", v.Name, v.Field)
	}

	e = v.Validate(
		&vv.StringIsAlpha{
			Name:  "Test",
			Field: "N0t Alpha",
		},
	)

	fmt.Println("After change: ", e.JSON())
}
