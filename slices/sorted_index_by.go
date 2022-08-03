package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like SortedIndex except that it accepts a function
// which is invoked for value and each element of slice to compute
// their sort ranking. The function is invoked with one argument: (value).
//
// example for 'function':
//
//	func makeInt(input interface{}) interface{} {
//	  return int(input.(float64))
//	}
//
// Complexity: O(log(n))
func SortedIndexBy(slice, value interface{}, function func(interface{}) interface{}) int {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	sType := reflect.TypeOf(slice)
	val := reflect.ValueOf(value)
	if val.Type().Kind() != sType.Elem().Kind() && sType.Elem().Kind() != reflect.Interface {
		panic("'value' is not compatible with 'slice' elements")
	}

	return whereToPutInSliceLowerEqualBy(slice, value, compareLowerEqual, function)
}

// Based on binary search, searchs on where to put the
// sent value in the passed slice based on isLowerEqualFunction
// result and get the comparators by getComparatorParam function
//
// Complexity: O(log(n))
func whereToPutInSliceLowerEqualBy(slice, value, isLowerEqualFunction, comparableParamFunction interface{}) int {
	sliceValue := reflect.ValueOf(slice)
	comparableParamFunctionValue := reflect.ValueOf(comparableParamFunction)
	len := sliceValue.Len()

	if len == 0 {
		return 0
	} else if len == 1 {
		item := sliceValue.Index(0)
		if res := reflect.ValueOf(isLowerEqualFunction).Call([]reflect.Value{item, reflect.ValueOf(value)}); res[0].Bool() {
			return 0
		} else {
			return 1
		}
	} else if len == 2 {
		item := sliceValue.Index(0)
		if res := reflect.ValueOf(isLowerEqualFunction).Call([]reflect.Value{item, reflect.ValueOf(value)}); res[0].Bool() {
			return 0
		} else {
			var result int
			if result = whereToPutInSliceLowerEqual(sliceValue.Slice(1, 2).Interface(), value, isLowerEqualFunction); result == -1 {
				return -1
			}
			return result + 1
		}
	}

	item := sliceValue.Index(len / 2).Interface()

	if !internal.AreComparable(item, value) {
		panic("couldn't compare 'value' with all items in passed slice")
	}

	var result int
	firstItem := comparableParamFunctionValue.Call([]reflect.Value{reflect.ValueOf(item)})[0]
	secondItem := comparableParamFunctionValue.Call([]reflect.Value{reflect.ValueOf(value)})[0]
	if res := reflect.ValueOf(isLowerEqualFunction).Call([]reflect.Value{firstItem, secondItem}); res[0].Bool() {
		if result = whereToPutInSliceLowerEqualBy(sliceValue.Slice(0, (len/2)+1).Interface(), value, isLowerEqualFunction, comparableParamFunction); result == -1 {
			return -1
		}

		return result
	} else {
		if result = whereToPutInSliceLowerEqualBy(sliceValue.Slice(len/2, len).Interface(), value, isLowerEqualFunction, comparableParamFunction); result == -1 {
			return -1
		}

		return result + (len / 2)
	}
}
