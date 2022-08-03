package slices

import (
	"fmt"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Converts all elements in slice into a string separated by 'separator'.
//
// Complexity: O(n)
func Join(slice interface{}, separator string) string {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)
	var result string
	for i := 0; i < sliceValue.Len(); i++ {
		if len(result) == 0 {
			result = toString(sliceValue.Index(0).Interface())
		} else {
			result += separator + toString(sliceValue.Index(i).Interface())
		}
	}
	return result
}

func toString(input interface{}) string {
	inputValue := reflect.ValueOf(input)
	if !inputValue.IsValid() {
		return ""
	}

	return fmt.Sprint(input)
}
