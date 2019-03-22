package main

import (
	"fmt"
	"strings"

	v "github.com/s3rj1k/validator"
)

/*
Simple validator:

type Validator interface {
	Validate(e *v.Errors)
}
*/

// StringHasLove is a Validator object
type StringHasLove struct {
	Name string // Mandatory Field

	// Amount/types of other fields are not limited
	Field string
}

// Validate adds an error if Field has no love.
func (cv *StringHasLove) Validate(e *v.Errors) {

	if strings.Contains(strings.ToLower(cv.Field), "love") {
		return
	}

	e.Add(cv.Name, fmt.Sprintf("'%s' has no love", cv.Field))
}

func main() {

	s := "I love donuts"

	e := v.Validate(
		&StringHasLove{
			Name:  "donuts",
			Field: s,
		},
	)
	fmt.Println("donuts errors: ", e.JSON())

	s = "I see dead people"
	e = v.Validate(
		&StringHasLove{
			Name:  "dead",
			Field: s,
		},
	)
	fmt.Println("dead errors: ", e.JSON())
}
