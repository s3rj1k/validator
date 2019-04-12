package main

import (
	"fmt"
	"regexp"
	"strings"

	v "github.com/s3rj1k/validator"
	vv "github.com/s3rj1k/validator/buildin"
)

/*
NumberValidator interface is used for NumberSliceDive.
Implement all the methods on any struct to pass it to NumberSliceDive.

type NumberValidator interface {
	Validate(*validator.Errors)
	SetField(interface{})
	SetNameIndex(int)
	GetName() string
}
*/

// NumberIsDivisibleBy is a Validator object
type NumberIsDivisibleBy struct {
	Name string // Field is mandatory

	// Amount of other fields are not limited
	Field         int
	ComparedField int
}

// Validate adds an error if Field has no love.
func (cv *NumberIsDivisibleBy) Validate(e *v.Errors) {

	// add your logic
	if cv.Field%cv.ComparedField == 0 {
		return
	}

	// add error
	e.Add(cv.Name, fmt.Sprintf("'%d' is not divisible by '%d'", cv.Field, cv.ComparedField))
}

// SetField sets validator field.
func (cv *NumberIsDivisibleBy) SetField(s interface{}) {
	cv.Field = s.(int) // change this casting to your type
}

// SetNameIndex sets index of slice element on Name.
// Do not change this method.
func (cv *NumberIsDivisibleBy) SetNameIndex(i int) {
	cv.Name = fmt.Sprintf("%s[%d]", rxSetNameIndex.ReplaceAllString(cv.Name, ""), i)
}

// GetName is a getter on Name field.
// Do not change this method.
func (cv *NumberIsDivisibleBy) GetName() string {
	return cv.Name
}

func main() {

	slice := []int{10, 33, 44, 57, 99, 100, 0, -3, -4}
	divisor := 3

	e := v.Validate(
		&vv.NumberSliceDive{
			Validator: &NumberIsDivisibleBy{
				Name:          "slice",
				ComparedField: divisor,
			},
			Field: slice,
		},
		&vv.NumberSliceDive{ // use other build in
			Validator: &vv.NumberInRange{
				Name: "slice",
				Min:  -100,
				Max:  90,
			},
			Field: slice,
		},
		&vv.SliceIsUnique{ // not only numbers validators
			Name:  "slice",
			Field: slice,
		},
	)
	if e != nil {
		// to map
		emap := e.Lookup("slice")

		for k, v := range emap {
			fmt.Printf("key: %s, value: [%s]\n", k, strings.Join(v, ", "))
		}
	}
}
