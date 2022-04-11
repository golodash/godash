package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

func FindIndex(slice interface{}, function interface{}) (int, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return -1, err
	}

	functionType := reflect.TypeOf(function)
	if functionType.Kind() != reflect.Func {
		return -1, errors.New("`function` has to be function type")
	}
	if functionType.NumIn() != 1 {
		return -1, errors.New("`function` function inputs has to have just 1 input")
	}
	if functionType.In(0).Kind() != reflect.TypeOf(slice).Elem().Kind() &&
		functionType.In(0).Kind() != reflect.Interface {
		return -1, errors.New("`function` function input has to be the same type as `slice` variable elements or have to be `interface` type")
	}
	if functionType.NumOut() != 1 || functionType.Out(0).Kind() != reflect.Bool {
		return -1, errors.New("`function` function output has to be `bool` type and it has to have just 1 output")
	}

	sliceValue := reflect.ValueOf(slice)
	functionValue := reflect.ValueOf(function)
	for i := 0; i < sliceValue.Len(); i++ {
		if functionValue.Call([]reflect.Value{sliceValue.Index(i)})[0].Bool() {
			return i, nil
		}
	}

	return -1, nil
}
