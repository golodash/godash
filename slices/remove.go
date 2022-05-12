package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Removes all elements from slice that the passed function returns true on them
// and returns a slice of remaining elements and a slice of removed elements.
// The passed function will invoke with one argument.
func Remove(slice, function interface{}) (interface{}, interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, nil, err
	}

	functionType := reflect.TypeOf(function)
	sliceType := reflect.TypeOf(slice)
	if functionType.Kind() != reflect.Func {
		return nil, nil, errors.New("`function` input should be a function type")
	}
	if functionType.NumIn() != 1 || (sliceType.Elem().String() != functionType.In(0).String() && functionType.In(0).Kind() != reflect.Interface) {
		return nil, nil, errors.New("`function` should have a single input and be compatible with slice elements")
	}
	if functionType.NumOut() != 1 || functionType.Out(0).Kind() != reflect.Bool {
		return nil, nil, errors.New("`function` should have a single output and it has to be boolean type")
	}

	newSlice, err := internal.InterfaceToSlice(slice)
	if err != nil {
		return nil, nil, err
	}

	functionValue := reflect.ValueOf(function)
	removed := []interface{}{}
	for i := len(newSlice) - 1; i >= 0; i-- {
		if res := functionValue.Call([]reflect.Value{reflect.ValueOf(newSlice[i])}); res[0].Bool() {
			removed = append(removed, newSlice[i])
			newSlice = append(newSlice[0:i], newSlice[i+1:]...)
		}
	}

	return newSlice, removed, nil
}
