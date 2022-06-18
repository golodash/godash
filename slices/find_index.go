package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method finds first index of the given slice which the given
// function on that element, returns true.
//
// Complexity: O(n)
func FindIndex(slice, value interface{}) (int, error) {
	if ok := internal.SliceCheck(slice); !ok {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	for i := 0; i < sliceValue.Len(); i++ {
		if ok, _ := internal.Same(reflect.ValueOf(sliceValue.Index(i).Interface()).Interface(), value); ok {
			return i, nil
		}
	}

	return -1, nil
}
