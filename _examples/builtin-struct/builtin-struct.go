package main

import (
	"fmt"

	v "github.com/s3rj1k/validator"
	vv "github.com/s3rj1k/validator/validators"
)

type user struct {
	id        int
	firstName string
	lastName  string
	age       int
	email     string
}

func main() {

	u := &user{
		firstName: "John",
		lastName:  "doe",
		age:       2,
		email:     "john_doe@no.one",
	}

	err := structValidation(u)
	fmt.Println(err)
	// Output: {"user.age":["'2' not in range(12, 112)"],"user.id":["'user.id' must not be equal to 0"],"user.lastName":["'doe' must contain at least 1 uppercase"]}

	err = structValidationP(u) //
	fmt.Println(err)
	// Output: {"user":{"age":["'2' not in range(12, 112)"],"id":["'user.id' must not be equal to 0"],"lastName":["'doe' must contain at least 1 uppercase"]}}
	// notice output error structure

	// cast to get access to custom error method
	e := err.(*v.Errors)
	fmt.Println(e.Count())
	// Output: 3
}

func structValidation(u *user) error {
	return v.Validate(
		&vv.NumberIsNotZero{ // adds an error
			Name:  "user.id",
			Field: u.id,
		},
		&vv.StringIsAlpha{
			Name:  "user.firstName",
			Field: u.firstName,
		},
		&vv.StringIsPresent{
			Name:  "user.lastName",
			Field: u.lastName,
		},
		&vv.StringHasUpperCase{ // adds an error
			Name:  "user.lastName",
			Field: u.lastName,
		},
		&vv.NumberInRange{ // adds an error
			Name:  "user.age",
			Field: u.age,
			Min:   12,
			Max:   112,
		},
		&vv.StringIsEmail{
			Name:  "user.email",
			Field: u.email,
		},
	)
}

func structValidationP(u *user) error {
	return v.ValidateP(
		&vv.NumberIsNotZero{ // adds an error
			Name:  "user.id",
			Field: u.id,
		},
		&vv.StringIsAlpha{
			Name:  "user.firstName",
			Field: u.firstName,
		},
		&vv.StringIsPresent{
			Name:  "user.lastName",
			Field: u.lastName,
		},
		&vv.StringHasUpperCase{ // adds an error
			Name:  "user.lastName",
			Field: u.lastName,
		},
		&vv.NumberInRange{ // adds an error
			Name:  "user.age",
			Field: u.age,
			Min:   12,
			Max:   112,
		},
		&vv.StringIsEmail{
			Name:  "user.email",
			Field: u.email,
		},
	)
}
