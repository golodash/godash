package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Gets the (n)th element of a slice. If n is negative, the (nth)
// element from the end is returned.
//
// Complexity: O(1)
func Nth(slice interface{}, index int) interface{} {
	if ok := internal.SliceCheck(slice); !ok {
		panic("passed 'slice' variable is not slice type")
	}

	values := reflect.ValueOf(slice)
	if index < 0 {
		index = values.Len() + index
	}

	if index >= values.Len() {
		panic("index out of range")
	}

	if values.Len() == 0 {
		return nil
	}

	return values.Index(index).Interface()
}
