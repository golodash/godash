package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Computes the maximum value of slice.
//
// Complexity: O(n)
func MinBy(slice interface{}, function func(interface{}) interface{}) interface{} {
	if !internal.SliceCheck(slice) {
		panic("'slice' is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)

	if sliceValue.Len() == 0 {
		return nil
	}

	biggest := reflect.ValueOf(function(sliceValue.Index(0).Interface()))
	for i := 0; i < sliceValue.Len(); i++ {
		element := reflect.ValueOf(function(sliceValue.Index(i).Interface()))
		if res := internal.CompareNumbers(element.Interface(), biggest.Interface()); res == internal.Lower {
			biggest = element
		}
	}

	return biggest.Interface()
}
