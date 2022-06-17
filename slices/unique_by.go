package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a duplicate-free version of a slice that only
// keeps the first occurrence of each element.
//
// The order of result values is determined by the order
// they occur in the slice.
//
// Complexity: O(n)
//
// n = length of passed slice
func UniqueBy(slice interface{}, function func(interface{}) interface{}) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	sliceValue := reflect.ValueOf(slice)
	outputValue := reflect.MakeSlice(reflect.TypeOf(slice), 0, sliceValue.Len())
	m := map[interface{}]bool{}
	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i).Interface()
		compareItem := function(item)
		if _, ok := m[compareItem]; !ok {
			m[compareItem] = true
			outputValue = reflect.Append(outputValue, reflect.ValueOf(item))
		}
	}

	return outputValue.Interface(), nil
}
