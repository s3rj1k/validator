package main

import (
	"fmt"
	"strings"

	v "github.com/s3rj1k/validator"
	vv "github.com/s3rj1k/validator/buildin"
)

/*
StringValidator interface is used for StringSliceDive.
Implement all the methods on any struct to pass it to StringSliceDive.

type StringValidator interface {
	Validate(*v.Errors)
	SetField(s string)
	SetNameIndex(i int)
}
*/

// StringHasLove is a Validator object
type StringHasLove struct {
	Name string // Field is mandatory

	// Amount/types of other fields are not limited
	Field string
}

// Validate adds an error if Field has no love.
func (cv *StringHasLove) Validate(e *v.Errors) {

	// add any logic that you need
	if strings.Contains(strings.ToLower(cv.Field), "love") {
		return
	}

	// add error
	e.Add(cv.Name, fmt.Sprintf("'%s' has no love", cv.Field))
}

// SetField sets validator field.
// Do not change this method.
func (cv *StringHasLove) SetField(s string) {
	cv.Field = s
}

// SetNameIndex sets index of slice element on Name.
// Do not change this method.
func (cv *StringHasLove) SetNameIndex(i int) {
	cv.Name = fmt.Sprintf("%s[%d]", vv.RxSetNameIndex.ReplaceAllString(cv.Name, ""), i)
}

func main() {

	s := []string{"LOVE", "I do not know what love is", "But still have love", "I do not", ""}

	e := v.Validate(
		&vv.StringSliceDive{ // 2 errors for 3 and 4 indexes
			Validator: &StringHasLove{
				Name: "slice",
			},
			Field: s,
		},
		&vv.StringSliceDive{ // other validators can be added
			Validator: &vv.StringIsASCII{
				Name: "slice",
			},
			Field: s,
		},
		&vv.SliceIsNotEmpty{ // not necessary dive
			Name:  "slice",
			Field: s,
		},
	)
	if e != nil {
		fmt.Println(e) // Output: {"slice[3]":["'I do not' has no love"],"slice[4]":["'' has no love"]}
	}
}
