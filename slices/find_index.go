package slices

import (
	"reflect"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

// This method finds first index of the given slice which the given
// value equals on that element's value.
//
// Complexity: O(n)
func FindIndex(slice, value interface{}) int {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	for i := 0; i < sliceValue.Len(); i++ {
		if generals.Same(reflect.ValueOf(sliceValue.Index(i).Interface()).Interface(), value) {
			return i
		}
	}

	return -1
}
