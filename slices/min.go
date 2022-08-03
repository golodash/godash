package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Computes the maximum value of slice.
//
// Complexity: O(n)
func Min(slice interface{}) interface{} {
	if !internal.SliceCheck(slice) {
		panic("'slice' is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)

	if sliceValue.Len() == 0 {
		return nil
	}

	biggest := sliceValue.Index(0)
	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i)
		if res := internal.CompareNumbers(element.Interface(), biggest.Interface()); res == internal.Lower {
			biggest = element
		}
	}

	return biggest.Interface()
}
