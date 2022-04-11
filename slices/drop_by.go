package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a new slice from the passed slice and removes elements
// from it when passed function returns true on that element.
//
// By default n is equal to 0
func DropBy(slice interface{}, function interface{}) ([]interface{}, error) {
	s, err := internal.InterfaceToSlice(slice)
	if err != nil {
		return nil, err
	}

	if reflect.TypeOf(function).Kind() != reflect.Func {
		return nil, errors.New("`function` variable is not a function")
	}

	functionType := reflect.TypeOf(function)
	if functionType.Kind() != reflect.Func {
		return nil, errors.New("`function` variable has to be function type")
	}
	if functionType.NumIn() != 1 {
		return nil, errors.New("`function` inputs has to have just 1 input")
	}
	if functionType.In(0).Kind() != reflect.TypeOf(slice).Elem().Kind() && functionType.In(0).Kind() != reflect.Interface {
		return nil, errors.New("`function` input has to be the same type as `slice` and `notIncluded` elements or have to be `interface` type")
	}
	if functionType.NumOut() != 1 || functionType.Out(0).Kind() != reflect.Bool {
		return nil, errors.New("`function` output has to be `bool` type and it has to have just 1 output")
	}

	functionValue := reflect.ValueOf(function)
	for i := 0; i < len(s); i++ {
		res := functionValue.Call([]reflect.Value{reflect.ValueOf(s[i])})
		if res[0].Bool() {
			if i != 0 && i+1 < len(s) {
				s = append(s[0:i], s[i+1:]...)
			} else if i == 0 {
				s = s[i+1:]
			} else if i+1 >= len(s) {
				s = s[0:i]
			}
		}
	}

	return s, nil
}
