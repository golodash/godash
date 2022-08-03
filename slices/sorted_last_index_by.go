package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Uses a binary search to determine the Highest index at which value should be
// inserted into slice in order to maintain its sort order.
//
// example for 'function':
//
//	func makeInt(input interface{}) interface{} {
//	  return int(input.(float64))
//	}
//
// Complexity: O(log(n))
func SortedLastIndexBy(slice, value interface{}, function func(interface{}) interface{}) int {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	sType := reflect.TypeOf(slice)
	val := reflect.ValueOf(value)
	if val.Type().Kind() != sType.Elem().Kind() && sType.Elem().Kind() != reflect.Interface {
		panic("'value' is not compatible with 'slice' elements")
	}

	return whereToPutInSliceBiggerEqualBy(slice, value, compareHigherEqual, function)
}

// Based on binary search, searchs on where to put the
// sent value in the passed slice based on isBiggerEqualFunction
// result and get the comparators by getComparatorParam function
//
// Complexity: O(log(n))
func whereToPutInSliceBiggerEqualBy(slice, value, isBiggerEqualFunction, comparableParamFunction interface{}) int {
	sliceValue := reflect.ValueOf(slice)
	comparableParamFunctionValue := reflect.ValueOf(comparableParamFunction)
	len := sliceValue.Len()

	secondItem := comparableParamFunctionValue.Call([]reflect.Value{reflect.ValueOf(value)})[0]
	if len == 0 {
		return 0
	} else if len == 1 {
		item := comparableParamFunctionValue.Call([]reflect.Value{sliceValue.Index(0)})[0]
		if res := reflect.ValueOf(isBiggerEqualFunction).Call([]reflect.Value{item, secondItem}); res[0].Bool() {
			return 1
		} else {
			return 0
		}
	} else if len == 2 {
		item := comparableParamFunctionValue.Call([]reflect.Value{sliceValue.Index(1)})[0]
		if res := reflect.ValueOf(isBiggerEqualFunction).Call([]reflect.Value{item, secondItem}); res[0].Bool() {
			return 2
		} else {
			var result int
			if result = whereToPutInSliceBiggerEqual(sliceValue.Slice(0, 1).Interface(), value, isBiggerEqualFunction); result == -1 {
				return -1
			}
			return result
		}
	}

	item := sliceValue.Index(len / 2).Interface()

	if !internal.AreComparable(item, value) {
		panic("couldn't compare 'value' with all items in passed slice")
	}

	var result int
	firstItem := comparableParamFunctionValue.Call([]reflect.Value{reflect.ValueOf(item)})[0]
	if res := reflect.ValueOf(isBiggerEqualFunction).Call([]reflect.Value{firstItem, secondItem}); res[0].Bool() {
		if result = whereToPutInSliceBiggerEqualBy(sliceValue.Slice(len/2, len).Interface(), value, isBiggerEqualFunction, comparableParamFunction); result == -1 {
			return -1
		}

		return result + (len / 2)
	} else {
		if result = whereToPutInSliceBiggerEqualBy(sliceValue.Slice(0, (len/2)+1).Interface(), value, isBiggerEqualFunction, comparableParamFunction); result == -1 {
			return -1
		}

		return result
	}
}
