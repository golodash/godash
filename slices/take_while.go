package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a sub slice of a slice with elements taken from the beginning.
// Elements are taken until passed function returns false.
//
// Complexity: O(n)
//
// n = number of elements that passed function returns true on them
func TakeWhile(slice interface{}, function func(interface{}) bool) interface{} {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	i := 0
	for i = 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i).Interface()
		if !function(item) {
			break
		}
	}

	return sliceValue.Slice(i, sliceValue.Len()).Interface()
}
