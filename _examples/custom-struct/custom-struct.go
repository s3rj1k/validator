package main

import (
	"fmt"
	"time"

	v "github.com/s3rj1k/validator"
)

/*
This example shows how to use the validator if you already have
a struct that needs validation
*/

// MyStruct is an existing struct
type MyStruct struct {
	ID      uint
	URL     string
	Timeout time.Duration
	Body    string
}

/*
To use the validator you have to define a Validate method
*/

// Validate performs struct validaton
func (cv *MyStruct) Validate(e *v.Errors) {

	// validating that ID > 0 and URL is not empty
	if cv.ID > 0 && cv.URL != "" {
		return
	}

	// Adding error if any of the conditions have failed.
	e.Add("MyStruct", fmt.Sprintf("MyStruct validation failed"))
	// First argument to e.Add() method is the path where errors of validation will be stored.
	// Idiomatic way is to define Name field in the struct (see other examples).
	// Another option is to hard-code the path (done here). In this case the DIVE is impossible.
	// Yet another options are to use any other string field / global variable as the path.
}

func main() {

	ms := &MyStruct{ // no ID defined
		URL: "www.google.com",
	}

	e := v.Validate(ms)
	if e != nil {
		fmt.Println(e) // Output: {"MyStruct":["MyStruct validation failed"]}
	}
}
