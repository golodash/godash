package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like 'difference' except that it accepts a custom function
// which is invoked to compare elements of 'slice' to 'notIncluded' slice.
//
// function has to indicate if two given variables as inputs are equal or not.
//
// example for 'function':
//
// 	func isEqual(value1, value2 interface{}) bool {
// 		return value1.(int) == value2.(int)
// 	}
//
// Complexity: O(n)
func DifferenceBy(slice, notIncluded interface{}, function func(interface{}, interface{}) bool) (interface{}, error) {
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
			if ok := function(reflect.ValueOf(sliceValue.Index(i).Interface()).Interface(), reflect.ValueOf(notInValue.Index(j).Interface()).Interface()); ok {
				sliceValue = reflect.AppendSlice(sliceValue.Slice(0, i), sliceValue.Slice(i+1, sliceValue.Len()))
				i++
				break firstLoop
			}
		}
	}

	return sliceValue.Interface(), nil
}
