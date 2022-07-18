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
// Complexity: O(n*m)
//
// n = length of 'slice'
//
// m = length of 'notIncluded'
func DifferenceBy(slice, notIncluded interface{}, function func(interface{}, interface{}) bool) interface{} {
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
			if ok := function(reflect.ValueOf(sliceValue.Index(i).Interface()).Interface(), reflect.ValueOf(notInValue.Index(j).Interface()).Interface()); ok {
				sliceValue = reflect.AppendSlice(sliceValue.Slice(0, i), sliceValue.Slice(i+1, sliceValue.Len()))
				i++
				break firstLoop
			}
		}
	}

	return sliceValue.Interface()
}
