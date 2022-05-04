package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Removes all elements from slice that the passed function returns true on them
// and returns a slice of remaining elements and a slice of removed elements.
// The passed function will invoke with one argument.
func Remove(slice interface{}, function interface{}) (interface{}, interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, nil, err
	}
	funcType := reflect.TypeOf(function)
	if funcType.Kind() != reflect.Func {
		return nil, nil, errors.New("the second input should be a function")
	}
	if funcType.NumIn() != 1 {
		return nil, nil, errors.New("function should have a single input")
	}
	if funcType.NumOut() != 1 || funcType.Out(0).Kind() != reflect.Bool {
		return nil, nil, errors.New("fnuction's output should be a single output and boolean type")
	}
	newSlice, err := internal.InterfaceToSlice(slice)
	if err != nil {
		return nil, nil, err
	}
	funcValues := reflect.ValueOf(function)
	var removed []interface{}

	for i := len(newSlice) - 1; i >= 0; i-- {
		res := funcValues.Call([]reflect.Value{reflect.ValueOf(newSlice[i])})
		if res[0].Bool() {
			removed = append(removed, newSlice[i])
			if i != 0 && i+1 < len(newSlice) {
				newSlice = append(newSlice[0:i], newSlice[i+1:]...)
			} else if i == 0 {
				newSlice = newSlice[i+1:]
			} else if i+1 == len(newSlice) {
				newSlice = newSlice[0:i]
			}
		}
	}
	return newSlice, removed, nil
}
