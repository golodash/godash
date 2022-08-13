package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like Min except that it accepts a function which
// is invoked for each element in slice to return a number for comparison.
//
// example for 'function':
//
//	type myObject struct {
//		rank int
//	}
//
//	func returnRank(value1 interface{}) interface{} {
//		return value1.(myObject).rank
//	}
//
// Complexity: O(n)
func MinBy(slice interface{}, function func(interface{}) interface{}) interface{} {
	if !internal.SliceCheck(slice) {
		panic("'slice' is not slice type")
	}

	sliceValue := reflect.ValueOf(slice)

	if sliceValue.Len() == 0 {
		return nil
	}

	biggest := reflect.ValueOf(function(sliceValue.Index(0).Interface())).Interface()
	chosenElement := sliceValue.Index(0)
	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i)
		compareValue := reflect.ValueOf(function(element.Interface())).Interface()
		if res := internal.CompareNumbers(compareValue, biggest); res == internal.Lower {
			biggest = compareValue
			chosenElement = element
		}
	}

	return chosenElement.Interface()
}
