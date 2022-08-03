package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Returns a slice of 'slice' elements that are not included in the
// other given slice using comparisons.
//
// Complexity: O(n*m)
//
// n = length of 'slice'
//
// m = length of 'notIncluded'
func Difference(slice, notIncluded interface{}) interface{} {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}
	if !internal.SliceCheck(notIncluded) {
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
			if internal.Same(sliceValue.Index(i).Interface(), notInValue.Index(j).Interface()) {
				sliceValue = reflect.AppendSlice(sliceValue.Slice(0, i), sliceValue.Slice(i+1, sliceValue.Len()))
				i++
				break firstLoop
			}
		}
	}

	return sliceValue.Interface()
}
