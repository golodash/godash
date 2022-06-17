package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Uses a binary search to determine the Highest index at which value should be
// inserted into slice in order to maintain its sort order.
//
// example for 'function':
//
//  func makeInt(input interface{}) interface{} {
//    return int(input.(float64))
//  }
//
// Complexity: O(log(n))
func SortedLastIndexBy(slice, value interface{}, function func(interface{}) interface{}) (int, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return -1, err
	}

	sType := reflect.TypeOf(slice)
	val := reflect.ValueOf(value)
	if val.Type().Kind() != sType.Elem().Kind() && sType.Elem().Kind() != reflect.Interface {
		return -1, errors.New("'value' is not compatible with 'slice' elements")
	}

	return whereToPutInSliceBiggerEqualBy(slice, value, compareHigherEqual, function)
}

// Based on binary search, searchs on where to put the
// sent value in the passed slice based on isBiggerEqualFunction
// result and get the comparators by getComparatorParam function
//
// Complexity: O(log(n))
func whereToPutInSliceBiggerEqualBy(slice, value, isBiggerEqualFunction, comparableParamFunction interface{}) (int, error) {
	sliceValue := reflect.ValueOf(slice)
	comparableParamFunctionValue := reflect.ValueOf(comparableParamFunction)
	len := sliceValue.Len()

	secondItem := comparableParamFunctionValue.Call([]reflect.Value{reflect.ValueOf(value)})[0]
	if len == 0 {
		return 0, nil
	} else if len == 1 {
		item := comparableParamFunctionValue.Call([]reflect.Value{sliceValue.Index(0)})[0]
		if res := reflect.ValueOf(isBiggerEqualFunction).Call([]reflect.Value{item, secondItem}); res[0].Bool() {
			return 1, nil
		} else {
			return 0, nil
		}
	} else if len == 2 {
		item := comparableParamFunctionValue.Call([]reflect.Value{sliceValue.Index(1)})[0]
		if res := reflect.ValueOf(isBiggerEqualFunction).Call([]reflect.Value{item, secondItem}); res[0].Bool() {
			return 2, nil
		} else {
			var result int
			var err error
			if result, err = whereToPutInSliceBiggerEqual(sliceValue.Slice(0, 1).Interface(), value, isBiggerEqualFunction); err != nil {
				return -1, err
			}
			return result, nil
		}
	}

	item := sliceValue.Index(len / 2).Interface()

	var err error = nil
	if err = internal.AreComparable(item, value); err != nil {
		return -1, errors.New("couldn't compare 'value' with all items in passed slice")
	}

	var result int
	firstItem := comparableParamFunctionValue.Call([]reflect.Value{reflect.ValueOf(item)})[0]
	if res := reflect.ValueOf(isBiggerEqualFunction).Call([]reflect.Value{firstItem, secondItem}); res[0].Bool() {
		if result, err = whereToPutInSliceBiggerEqualBy(sliceValue.Slice(len/2, len).Interface(), value, isBiggerEqualFunction, comparableParamFunction); err != nil {
			return -1, err
		}

		return result + (len / 2), nil
	} else {
		if result, err = whereToPutInSliceBiggerEqualBy(sliceValue.Slice(0, (len/2)+1).Interface(), value, isBiggerEqualFunction, comparableParamFunction); err != nil {
			return -1, err
		}

		return result, nil
	}
}
