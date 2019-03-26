# github.com/s3rj1k/validator

This package provides a framework for writing validations for Go data structures.

## Installation

```bash
$ go get github.com/s3rj1k/validator
```

## Usage

See examples below and in [_examples](https://github.com/s3rj1k/validator/tree/master/_examples) folder of this repo.


## Built-in Validators

This package has a children package with some nice built-in validators. 
*Number* validators support any *int* and *uint* types. *Slice* validators support all Go basic types except for bool.  


Follow rules and recommendations:
1. Validation can be done using *Validate* of *ValidateP* functions of the main package. The only difference is the structure of
output errors. See [builtin-struct example.](https://github.com/s3rj1k/validator/blob/master/_examples/builtin-struct/builtin-struct.go)
2. *Name* field **must** be passed to each built-in validator for proper validation.
3. Build-in validators have defaul error messages. They can be redefined, see [errmsgchange](https://github.com/s3rj1k/validator/tree/master/_examples/errmsgchange) example.
4. Note that validation errors is a custom type which has various methods for conveniece. See [godoc](https://godoc.org/github.com/s3rj1k/validator) for additional information.
5. Note that *nil* Field passed to Number validators will be converted to zero which can lead to unexpected validation effect.

```go
package main

import (
	"fmt"

	v "github.com/s3rj1k/validator"
	vv "github.com/s3rj1k/validator/validators"
)

type User struct {
	Name  string
	Email string
}

func main() {
	u := User{Name: "", Email: ""}
	e := v.Validate(
		&vv.StringIsPresent{
                        Name: "user.name", 
                        Field: u.Name
                        },
		&vv.StringIsEmail{
                        Name: "user.email", 
                        Field: u.Email, 
                        Message: "Mail is not in the right format." // custom error message
                        },
	)
	if e != nil {
		fmt.Println(e) // Output: {"user.email":["Mail is not in the right format."],"user.name":["'user.name' must not be blank"]}
	}
}
```


## Custom Validators

To validate a struct just define `Validate` method for a struct.
Below is a pretty simple example. Additional examples can be found in [_examples](https://github.com/s3rj1k/validator/tree/master/_examples) folder.

Follow rules and recommendations:
1. Add() method of v.Errors struct is the most important part where actual errors are added to the final validation results.
Make sure that e.Add() is called in Validate method.
2. For idiomatic usage of the validator - add a Name field to your struct (can have another name) and pass it as a first
argument to e.Add() method. The Name field is optional for simple validations (see [custom-struct example](https://github.com/s3rj1k/validator/blob/master/_examples/custom-struct/custom-struct.go)), but is mandatory
if you need to perform a dive (see any [dive examples](https://github.com/s3rj1k/validator/tree/master/_examples)). Omit the Name field ONLY if you are 100% sure in what you are doing.

```go
package main

import (
	"fmt"

	v "github.com/s3rj1k/validator"
)

type User struct {
	Name  string
	Email string
}

func (u *User) Validate(e *v.Errors) {
	if u.Name == "" {
		e.Add("name", "Name must not be blank!")
	}
	if u.Email == "" {
		e.Add("email", "Email must not be blank!")
	}
}

func main() {
	e := v.Validate(&User{})
	if e != nil {
		fmt.Println(e) // Output: {"email":["Email must not be blank!"],"name":["Name must not be blank!"]}
	}
}
```

