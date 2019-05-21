package main

import (
	"fmt"

	v "github.com/s3rj1k/validator"
	vv "github.com/s3rj1k/validator/buildin"
)

type user struct {
	firstName string
	age       int
}

func main() {
	u := &user{}

	// validate each field using separate validator structure
	e := v.Validate(
		&vv.StringIsPresent{ // check if string is not empty
			// Name is mandatory in each validator object.
			// Dot-notation is suggested since package is used for JSON-formatted output of errors.
			// Any format of Name field is OK while it is not empty.
			Name:  "user.firstName",
			Field: u.firstName,
		},
		&vv.NumberIsNotZero{ // check if number != 0
			Name:  "user.age",
			Field: u.age,
		},
	)
	fmt.Println(e)
	// Output: {"user.age":["'user.age' must not be equal to 0"],"user.firstName":["'user.firstName' must not be blank"]}
}
