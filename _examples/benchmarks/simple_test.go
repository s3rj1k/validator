package main

import (
	"testing"

	s3rj1k "github.com/s3rj1k/validator"
	vv "github.com/s3rj1k/validator/buildin"

	goplayg "gopkg.in/go-playground/validator.v9"

	asaskevich "github.com/asaskevich/govalidator"

	ozzo "github.com/go-ozzo/ozzo-validation"
)

type Test struct {
	Field1  string `validate:"required" valid:"required"` // `validate` for go-playground
	Field2  string `validate:"required" valid:"required"` // `valid` for asaskevich
	Field3  string `validate:"required" valid:"required"`
	Field4  string `validate:"required" valid:"required"`
	Field5  string `validate:"required" valid:"required"`
	Field6  int    `validate:"required" valid:"required"`
	Field7  int    `validate:"required" valid:"required"`
	Field8  uint   `validate:"required" valid:"required"`
	Field9  uint   `validate:"required" valid:"required"`
	Field10 uint   `validate:"required" valid:"required"`
}

var s = &Test{}

func Benchmark_s3rj1k_simple(b *testing.B) {
	message := "must not be null"
	for i := 0; i < b.N; i++ {
		err := s3rj1k.ValidateSync(
			&vv.StringIsPresent{
				Field:   s.Field1,
				Message: message,
			},
			&vv.StringIsPresent{
				Field:   s.Field2,
				Message: message,
			},
			&vv.StringIsPresent{
				Field:   s.Field3,
				Message: message,
			},
			&vv.StringIsPresent{
				Field:   s.Field4,
				Message: message,
			},
			&vv.StringIsPresent{
				Field:   s.Field5,
				Message: message,
			},
			&vv.NumberIsNotZero{
				Field:   s.Field6,
				Message: message,
			},
			&vv.NumberIsNotZero{
				Field:   s.Field7,
				Message: message,
			},
			&vv.NumberIsNotZero{
				Field:   s.Field8,
				Message: message,
			},
			&vv.NumberIsNotZero{
				Field:   s.Field9,
				Message: message,
			},
			&vv.NumberIsNotZero{
				Field:   s.Field10,
				Message: message,
			},
		)
		_panicnoerr(err)
	}
}

var goplg = goplayg.New()

func Benchmark_goplaygroud_simple(b *testing.B) {

	for i := 0; i < b.N; i++ {
		err := goplg.Struct(s)
		_panicnoerr(err)
	}
}

func Benchmark_asaskevich_simple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := asaskevich.ValidateStruct(s)
		_panicnoerr(err)
	}
}

func Benchmark_ozzo_simple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := ozzo.ValidateStruct(s,
			ozzo.Field(&s.Field1, ozzo.Required),
			ozzo.Field(&s.Field2, ozzo.Required),
			ozzo.Field(&s.Field3, ozzo.Required),
			ozzo.Field(&s.Field4, ozzo.Required),
			ozzo.Field(&s.Field5, ozzo.Required),
			ozzo.Field(&s.Field6, ozzo.Required),
			ozzo.Field(&s.Field7, ozzo.Required),
			ozzo.Field(&s.Field8, ozzo.Required),
			ozzo.Field(&s.Field9, ozzo.Required),
			ozzo.Field(&s.Field10, ozzo.Required),
		)
		_panicnoerr(err)
	}
}
