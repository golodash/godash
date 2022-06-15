package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like 'UniqueBy' except that it's designed and optimized for sorted slices.
//
// it accepts iteratee which is invoked for each element in array to generate the criterion by which uniqueness is computed.
func SortedUniqueBy(slice, function interface{}) (interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}

	sliceType := reflect.TypeOf(slice)
	funcType := reflect.TypeOf(function)
	if funcType.Kind() != reflect.Func {
		return nil, errors.New("'function' input should be function type")
	}
	if funcType.NumIn() != 1 || (funcType.In(0).String() != sliceType.Elem().String() && funcType.In(0).Kind() != reflect.Interface) {
		return nil, errors.New("'function' should have a single input and it's input type has to be compatible with slice elements types")
	}
	if funcType.NumOut() != 1 || !internal.IsNumberType(funcType.Out(0).Kind()) {
		return nil, errors.New("'function''s output should be a single output and it's type has to be a number")
	}

	newSlice, err := internal.InterfaceToSlice(slice)
	if err != nil {
		return nil, err
	}

	funcValue := reflect.ValueOf(function)
	m := make(map[interface{}]bool)
	unique := []interface{}{}
	for _, value := range newSlice {
		key := funcValue.Call([]reflect.Value{reflect.ValueOf(value)})[0].Interface()
		if _, ok := m[key]; !ok {
			m[key] = true
			unique = append(unique, value)
		}
	}

	return unique, nil
}
