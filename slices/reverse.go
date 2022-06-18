package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Reverses slice so that the first element becomes the last,
// the second element becomes the second to last, and so on.
//
// Complexity: O(n)
func Reverse(slice interface{}) (interface{}, error) {
	if ok := internal.SliceCheck(slice); !ok {
		panic("passed 'slice' variable is not slice type")
	}

	length := reflect.ValueOf(slice).Len()
	swapper := reflect.MakeSlice(reflect.TypeOf(slice), length, length).Interface()
	reflect.Copy(reflect.ValueOf(swapper), reflect.ValueOf(slice))
	swap := reflect.Swapper(swapper)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}

	return swapper, nil
}
