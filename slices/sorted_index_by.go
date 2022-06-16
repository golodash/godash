package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like SortedIndex except that it accepts a function
// which is invoked for value and each element of slice to compute
// their sort ranking. The function is invoked with one argument: (value).
//
// Complexity: O(log(n))
func SortedIndexBy(slice, value interface{}, function func(interface{}) interface{}) (int, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return -1, err
	}

	sType := reflect.TypeOf(slice)
	val := reflect.ValueOf(value)
	if val.Type().Kind() != sType.Elem().Kind() && sType.Elem().Kind() != reflect.Interface {
		return -1, errors.New("'value' is not compatible with 'slice' elements")
	}

	return whereToPutInSliceLowerEqualBy(slice, value, compareLowerEqual, function)
}

// Based on binary search, searchs on where to put the
// sent value in the passed slice based on isLowerEqualFunction
// result and get the comparators by getComparatorParam function
//
// Complexity: O(log(n))
func whereToPutInSliceLowerEqualBy(slice, value, isLowerEqualFunction, comparableParamFunction interface{}) (int, error) {
	sliceValue := reflect.ValueOf(slice)
	comparableParamFunctionValue := reflect.ValueOf(comparableParamFunction)
	len := sliceValue.Len()

	if len == 0 {
		return 0, nil
	} else if len == 1 {
		item := sliceValue.Index(0)
		if res := reflect.ValueOf(isLowerEqualFunction).Call([]reflect.Value{item, reflect.ValueOf(value)}); res[0].Bool() {
			return 0, nil
		} else {
			return 1, nil
		}
	} else if len == 2 {
		item := sliceValue.Index(0)
		if res := reflect.ValueOf(isLowerEqualFunction).Call([]reflect.Value{item, reflect.ValueOf(value)}); res[0].Bool() {
			return 0, nil
		} else {
			var result int
			var err error
			if result, err = whereToPutInSliceLowerEqual(sliceValue.Slice(1, 2).Interface(), value, isLowerEqualFunction); err != nil {
				return -1, err
			}
			return result + 1, nil
		}
	}

	item := sliceValue.Index(len / 2).Interface()

	var err error = nil
	if err = internal.AreComparable(item, value); err != nil {
		return -1, errors.New("couldn't compare 'value' with all items in passed slice")
	}

	var result int
	firstItem := comparableParamFunctionValue.Call([]reflect.Value{reflect.ValueOf(item)})[0]
	secondItem := comparableParamFunctionValue.Call([]reflect.Value{reflect.ValueOf(value)})[0]
	if res := reflect.ValueOf(isLowerEqualFunction).Call([]reflect.Value{firstItem, secondItem}); res[0].Bool() {
		if result, err = whereToPutInSliceLowerEqualBy(sliceValue.Slice(0, (len/2)+1).Interface(), value, isLowerEqualFunction, comparableParamFunction); err != nil {
			return -1, err
		}

		return result, nil
	} else {
		if result, err = whereToPutInSliceLowerEqualBy(sliceValue.Slice(len/2, len).Interface(), value, isLowerEqualFunction, comparableParamFunction); err != nil {
			return -1, err
		}

		return result + (len / 2), nil
	}
}
