package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Returns a slice of 'slice' elements that are not included in the
// other given slice using equality comparisons.
//
// Complexity: O(n*m)
//
// n = length of 'slice'
//
// m = length of 'notIncluded'
func Difference(slice, notIncluded interface{}) interface{} {
	if ok := internal.SliceCheck(slice); !ok {
		panic("passed 'slice' variable is not slice type")
	}
	if ok := internal.SliceCheck(notIncluded); !ok {
		panic("passed 'notIncluded' variable is not slice type")
	}

	notInValue := reflect.ValueOf(notIncluded)
	sliceValue := reflect.ValueOf(slice)
	for i := sliceValue.Len() - 1; i > -1; i-- {
		if i >= sliceValue.Len() {
			continue
		}
	firstLoop:
		for j := 0; j < notInValue.Len(); j++ {
			if ok := internal.Same(sliceValue.Index(i).Interface(), notInValue.Index(j).Interface()); ok {
				sliceValue = reflect.AppendSlice(sliceValue.Slice(0, i), sliceValue.Slice(i+1, sliceValue.Len()))
				i++
				break firstLoop
			}
		}
	}

	return sliceValue.Interface()
}

// Returns a slice of 'slice' elements that are not included in the
// other given slice using equality comparisons.
//
// Complexity: O(n*m)
//
// n = length of 'slice'
//
// m = length of 'notIncluded'
func Without(slice, notIncluded interface{}) interface{} {
	return Difference(slice, notIncluded)
}
