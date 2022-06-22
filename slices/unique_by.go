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
// example for 'function':
//
//  func makeInt(input interface{}) interface{} {
//    return int(input.(float64))
//  }
//
// Complexity: O(n)
func UniqueBy(slice interface{}, function func(interface{}) interface{}) interface{} {
	if ok := internal.SliceCheck(slice); !ok {
		panic("passed 'slice' variable is not slice type")
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

	return outputValue.Interface()
}
