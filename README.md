# github.com/s3rj1k/validator

This package provides a framework for writing validations for Go data structures.

## Installation

```bash
$ go get github.com/s3rj1k/validator
```

## General info

The package was developed with main purpose to provide neatly JSON-formatted errors of validation and idea to
limit usage of empty interfaces (and reflect package as a result).

The package provides 4 Validate functions:
*  ```Validate()``` - asynchronous function with simple JSON output.  Use this if you are not sure what you need.
*  ```ValidateP()``` - asynchronous function with dot notated JSON enabled. Use this if you need to embed error messages in JSON-output. 
*  ```ValidateSync()``` - synchronous function with simple JSON output. Use this if you need to increase performance of validations.
*  ```ValidateSyncP()``` - synchronous function with dot notated JSON enabled. Use this if you need to increase performance of validations and to embed error messages in JSON-output.
  
You can see 2 options of JSON-output of errors in examples below or [this example.](https://github.com/s3rj1k/validator/blob/master/_examples/builtin-struct/builtin-struct.go)
You can embed error message to any level of JSON output by using dot-notation in Name field of validator objects.



## Built-in Validators

This package has a children package with some nice built-in validators. Name of validator starts with the type of variable that needs to be validated (e.g. String, Number, Slice).  
*Number* validators support any *int* and *uint* types. *Slice* validators support all Go basic types except for bool.  


Follow rules and recommendations:
1. *Name* field **must** be passed to each built-in validator for proper validation.
2. Build-in validators have default error messages. They can be redefined, see [errmsgchange](https://github.com/s3rj1k/validator/tree/master/_examples/errmsgchange) example.
3. Note that validation errors is a custom type which has various methods for conveniece. See [godoc](https://godoc.org/github.com/s3rj1k/validator) for additional information.
4. Note that *nil* Field passed to Number validators will be converted to zero which can lead to unexpected validation effect.

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
			Name:  "user.name",
			Field: u.Name,
		},
		&vv.StringIsEmail{
			Name:    "user.email",
			Field:   u.Email,
			Message: "Mail is not in the right format.", // custom error message
		},
	)
	if e != nil {
		fmt.Println(e) // Output: {"user.email":["Mail is not in the right format."],"user.name":["'user.name' must not be blank"]}
	}
	// if ValidateP would have been used: {"user":{"email":["Mail is not in the right format."],"name":["'user.name' must not be blank"]}}
}

```
See other examples below and in [_examples](https://github.com/s3rj1k/validator/tree/master/_examples) folder of this repo.

## Custom Validators

To validate a struct just define `Validate` method for a struct. Below is a pretty simple example.

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

This, hovewer, is not an idiomatic way of this package to perform validation. Since first argument to e.Add() call is hardcoded, there is no
option to use JSON dot notation for output of errors (i.e. ValidateP output will be the same as Validate). If you do not care about JSON
formatting - example above is the way to go.

Below is a suggested example to do the same as in example above (but with ValidateP):

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

type StringIsNotBlank struct {
	Name string
	Field string
}

func (u *StringIsNotBlank) Validate(e *v.Errors) {
	if v.Field == "" {
		e.Add(v.Name, "Field must not be blank!")
	}
}

func main() {
	u := &User{}
	e := v.ValidateP(
		&StringIsNotBlank{
			Name: "user.name",
			Field: u.Name,
		},
		&StringIsNotBlank{
			Name: "user.email",
			Field: u.Email,
		},
	)
	if e != nil {
		fmt.Println(e) // Output: {"user":{"email":["Field must not be blank!"],"name":["Field must not be blank!"]}}
	}
}
```

Follow rules and recommendations:
1. Add() method of v.Errors struct is the most important part where actual errors are added to the final validation results.
Make sure that e.Add() is called in Validate method.
2. If you need to perform a slice dive see any [dive examples](https://github.com/s3rj1k/validator/tree/master/_examples).

## Benchmarks

[Benchmarks](https://github.com/s3rj1k/validator/tree/master/_examples/benchmarks) were written for cases that we expect to be the 
most relevant for the validator usage. I.e. struct with 5-10 fields each field requiring a validation.  
Simple benchmarks use simle checks on 10 fields of a structure. Increase in both memory and cpu is
achieved by not using empty interfaces/reflect package, but requires more code at Validate() call (compared to go-playground or asaskevich).  
Complex benchmarks include Get requests to existing and non-existing URLs. Dramatic increase in ns/op is achieved by 
asynchronous flow provided by Validate() func of the package.  

```
$ go test ./_examples/benchmarks -bench . -benchmem
goos: linux
goarch: amd64
pkg: github.com/s3rj1k/validator/_examples/benchmarks

Benchmark_s3rj1k_simple-4                 500000              2319 ns/op            1488 B/op         23 allocs/op
Benchmark_goplayground_simple-4           300000              4574 ns/op            2488 B/op         56 allocs/op
Benchmark_asaskevich_simple-4             100000             12796 ns/op            7491 B/op         86 allocs/op
Benchmark_ozzo_simple-4                   100000             13340 ns/op            8438 B/op        153 allocs/op

Benchmark_s3rj1k_complex-4                     5         296326162 ns/op          358256 B/op        882 allocs/op
Benchmark_goplayground_complex-4               2         577359003 ns/op          342840 B/op        890 allocs/op
Benchmark_ozzo_complex-4                       2         582399228 ns/op          335172 B/op        889 allocs/op
Benchmark_asaskevich_complex-4                 2         565452908 ns/op          405168 B/op        953 allocs/op
```