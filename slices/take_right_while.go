package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a sub slice of a slice with elements taken from the end.
// Elements are taken until passed function returns false.
//
// example for 'function':
//
//  func isEven(input interface{}) bool {
//    return input.(int) % 2 == 0
//  }
//
// Complexity: O(n)
//
// n = number of elements that passed function returns true on them
func TakeRightWhile(slice interface{}, function func(interface{}) bool) (interface{}, error) {
	if ok := internal.SliceCheck(slice); !ok {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	i := 0
	for i = sliceValue.Len() - 1; i > -1; i-- {
		if !function(sliceValue.Index(i).Interface()) {
			break
		}
	}

	return sliceValue.Slice(0, i+1).Interface(), nil
}
