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
func Difference(slice, notIncluded interface{}) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}
	if err := internal.SliceCheck(notIncluded); err != nil {
		return nil, err
	}

	notInValue := reflect.ValueOf(notIncluded)
	sliceValue := reflect.ValueOf(slice)
	for i := sliceValue.Len() - 1; i > -1; i-- {
		if i >= sliceValue.Len() {
			continue
		}
	firstLoop:
		for j := 0; j < notInValue.Len(); j++ {
			res, err := internal.Same(sliceValue.Index(i).Interface(), notInValue.Index(j).Interface())
			if err != nil {
				return nil, err
			}
			if res {
				sliceValue = reflect.AppendSlice(sliceValue.Slice(0, i), sliceValue.Slice(i+1, sliceValue.Len()))
				i++
				break firstLoop
			}
		}
	}

	return sliceValue.Interface(), nil
}

// Returns a slice of 'slice' elements that are not included in the
// other given slice using equality comparisons.
//
// Complexity: O(n*m)
//
// n = length of 'slice'
//
// m = length of 'notIncluded'
func Without(slice, notIncluded interface{}) (interface{}, error) {
	return Difference(slice, notIncluded)
}
