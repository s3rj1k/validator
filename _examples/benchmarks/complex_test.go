package main

import (
	"errors"
	"net/http"
	"net/url"
	"testing"
	"time"

	v "github.com/s3rj1k/validator"

	goplayg "gopkg.in/go-playground/validator.v9"

	asaskevich "github.com/asaskevich/govalidator"

	ozzo "github.com/go-ozzo/ozzo-validation"
)

var TestClient = http.Client{
	Timeout: time.Second * 5,
}

type Test2 struct {
	Field1 string `validate:"status200" valid:"status200"`
	Field2 string `validate:"status200" valid:"status200"`
	Field3 string `validate:"status200" valid:"status200"`
	Field4 string `validate:"status200" valid:"status200"`
	Field5 string `validate:"status200" valid:"status200"`
}

var test = &Test2{
	Field1: "http://www.google.com",
	Field2: "http://www.yahoo.com",
	Field3: "http://www.facebook.com",
	Field4: "http://www.thismustnotexistffs.prg",
	Field5: "http://www.ihopenotexistingurl.com",
}

func init() {
	// register for go-playground
	gopl.RegisterValidation("status200", ValidateStatus200)

	// register for asaskevich
	asaskevich.TagMap["status200"] = asaskevich.Validator(func(str string) bool {
		return statusCode200(str)
	})
}

// use a single instance of Validate, it caches struct info
var gopl = goplayg.New()

// statusCode200 wraps logic of complex validation
func statusCode200(s string) bool {
	parsed, err := url.Parse(s)
	if err != nil {
		panic("cant parse url")
	}

	resp, e := TestClient.Do(&http.Request{URL: parsed})
	if e != nil {
		return false
	}

	return resp.StatusCode == http.StatusOK
}

// StringIsStatusCode200 is a struct for s3rj1k validator
type StringIsStatusCode200 struct {
	Name  string
	Field string
}

// Validate is a method for struct of s3rj1k validator
func (v *StringIsStatusCode200) Validate(e *v.Errors) {

	if statusCode200(v.Field) {
		return
	}

	e.Add(v.Name, "status not 200")
}

// ValidateStatus200 is a func for go-playground validator
func ValidateStatus200(fl goplayg.FieldLevel) bool {
	return statusCode200(fl.Field().String())
}

// checkStatusCode200 is a func for ozzo validator
func checkStatusCode200(value interface{}) error {
	if statusCode200(value.(string)) {
		return errors.New("status code not 200")
	}
	return nil
}

func Benchmark_s3rj1k_complex(b *testing.B) {

	for i := 0; i < b.N; i++ {

		err := v.Validate(
			&StringIsStatusCode200{
				Field: test.Field1,
			},
			&StringIsStatusCode200{
				Field: test.Field2,
			},
			&StringIsStatusCode200{
				Field: test.Field3,
			},
			&StringIsStatusCode200{
				Field: test.Field4,
			},
			&StringIsStatusCode200{
				Field: test.Field5,
			},
		)
		_panicnoerr(err)

	}
}

func Benchmark_goplayground_complex(b *testing.B) {

	for i := 0; i < b.N; i++ {
		err := gopl.Struct(test)
		_panicnoerr(err)
	}
}

func Benchmark_ozzo_complex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ozzo.ValidateStruct(test,
			ozzo.Field(&test.Field1, ozzo.By(checkStatusCode200)),
			ozzo.Field(&test.Field2, ozzo.By(checkStatusCode200)),
			ozzo.Field(&test.Field3, ozzo.By(checkStatusCode200)),
			ozzo.Field(&test.Field4, ozzo.By(checkStatusCode200)),
			ozzo.Field(&test.Field5, ozzo.By(checkStatusCode200)),
		)
		_panicnoerr(err)
	}
}

func Benchmark_asaskevich_complex(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_, err := asaskevich.ValidateStruct(test)
		_panicnoerr(err)
	}
}

func _panicnoerr(e error) {
	if e == nil {
		panic("no errors")
	}
}
