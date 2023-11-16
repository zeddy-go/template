package repository

import (
	"reflect"
	"testing"
)

func TestXxx(t *testing.T) {
	ft := reflect.TypeOf(NewUserRepository)
	println(ft.Out(0).String())
}
