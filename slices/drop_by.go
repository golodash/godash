package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a new slice from the passed slice and removes elements
// from it when passed function returns true on that element.
//
// example for 'function':
//
//	func isEqual(input interface{}) bool {
//	  return input.(int)%2 == 0
//	}
//
// Complexity: O(n)
func DropBy(slice interface{}, function func(interface{}) bool) interface{} {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	for i := 0; i < sliceValue.Len(); i++ {
		res := function(sliceValue.Index(i).Interface())
		if res {
			sliceValue = reflect.AppendSlice(sliceValue.Slice(0, i), sliceValue.Slice(i+1, sliceValue.Len()))
		}
	}

	return sliceValue.Interface()
}
