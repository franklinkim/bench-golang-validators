package main

import (
	"testing"

	"github.com/go-playground/validator/v10"
	saddam "github.com/thedevsaddam/govalidator"
)

func BenchmarkAsaskevich(b *testing.B) {
	user := newUser()

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		validateWithAsaskevich(user)
	}
}

func BenchmarkPlayground(b *testing.B) {
	user := newUser()
	validate := validator.New()

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		validateWithPlayground(validate, user)
	}
}

func BenchmarkSaddam(b *testing.B) {
	user := newUser()

	rules := saddam.MapData{
		"firstname": []string{"required"},
		"lastname":  []string{"required"},
		"age":       []string{"min:0", "max:130"},
		"email":     []string{"required", "email"},
		"street":    []string{"required"},
		"city":      []string{"required"},
		"planet":    []string{"required"},
		"phone":     []string{"required"},
	}
	opts := saddam.Options{
		Data:  user,
		Rules: rules,
	}
	v := saddam.New(opts)

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		e := v.ValidateStruct()
		if len(e) > 0 {
			println(e)
		}
	}
}

func BenchmarkBuffalo(b *testing.B) {
	user := newUser()

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		validateWithBuffalo(user)
	}
}

func BenchmarkFender(b *testing.B) {
	user := newUser()

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		validateWithFender(user)
	}
}

func BenchmarkOzzo(b *testing.B) {
	user := newUser()

	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		validateWithOzzo(user)
	}
}
