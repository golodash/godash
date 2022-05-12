package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

//This method is like 'UniqueBy' except that it's designed and optimized for sorted slices.
//it accepts iteratee which is invoked for each element in array to generate the criterion by which uniqueness is computed.
func SortedUniqueBy(slice interface{}, function interface{}) (interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}
	funcType := reflect.TypeOf(function)
	if funcType.Kind() != reflect.Func {
		return nil, errors.New("the second input shoulbe a function")
	}
	if funcType.NumIn() != 1 {
		return nil, errors.New("function should have a single input")
	}
	if funcType.NumOut() != 1 {
		return nil, errors.New("fnuction's output should be a single output and boolean type")
	}

	var output []interface{}
	funcValues := reflect.ValueOf(function)
	newSlice, err := internal.InterfaceToSlice(slice)
	if err != nil {
		return nil, err
	}
	for i := len(newSlice) - 1; i >= 0; i-- {
		res := funcValues.Call([]reflect.Value{reflect.ValueOf(newSlice[i])})
		output = append(output, res[0].Interface())
	}
	unique, err := SortedUnique(output)
	if err != nil {
		return nil, err
	}

	return unique, nil
}
