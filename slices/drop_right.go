package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a sub slice from passed slice with n elements dropped from the end.
//
// Complexity: O(1)
func DropRight(slice interface{}, n int) interface{} {
	if ok := internal.SliceCheck(slice); !ok {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Len() < n {
		panic("'num' is bigger than slice length")
	}
	if n < 0 {
		panic("'num' is lower that zero")
	}

	return sliceValue.Slice(0, sliceValue.Len()-n).Interface()
}
