package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a duplicate-free version of a slice that only
// keeps the first occurrence of each element.
//
// The order of result values is determined by the order
// they occur in the slice.
func UniqueBy(slice, function interface{}) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	functionType := reflect.TypeOf(function)
	sliceType := reflect.TypeOf(slice)
	if functionType.Kind() != reflect.Func {
		return nil, errors.New("'function' input should be a function type")
	}
	if functionType.NumIn() != 1 || (sliceType.Elem().String() != functionType.In(0).String() && functionType.In(0).Kind() != reflect.Interface) {
		return nil, errors.New("'function' should have a single input and be compatible with slice elements")
	}
	if functionType.NumOut() != 1 {
		return nil, errors.New("'function' should have a single output")
	}

	functionValue := reflect.ValueOf(function)
	sliceValue := reflect.ValueOf(slice)
	outputValue := reflect.MakeSlice(reflect.TypeOf(slice), 0, sliceValue.Len())
	m := map[interface{}]bool{}
	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i).Interface()
		compareItem := functionValue.Call([]reflect.Value{reflect.ValueOf(item)})[0].Interface()
		if _, ok := m[compareItem]; !ok {
			m[compareItem] = true
			outputValue = reflect.Append(outputValue, reflect.ValueOf(item))
		}
	}

	return outputValue.Interface(), nil
}
