package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like LastIndexOf except that it performs a
// binary search on a sorted slice.
//
// Complexity: O(log(n))
func SortedLastIndexOf(slice, value interface{}) (int, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return -1, err
	}

	sType := reflect.TypeOf(slice)
	val := reflect.ValueOf(value)
	if val.Type().Kind() != sType.Elem().Kind() && sType.Elem().Kind() != reflect.Interface {
		return -1, errors.New("'value' is not compatible with 'slice' elements")
	}

	return sortedLastIndexOf(slice, value, compareHigherEqual, compareIsEqual)
}

// Complexity: O(log(n))
func sortedLastIndexOf(slice, value, isHigherEqualFunction, isEqualFunction interface{}) (int, error) {
	sliceValue := reflect.ValueOf(slice)
	len := sliceValue.Len()

	if len == 0 {
		return -1, errors.New("item not found")
	} else if len == 1 {
		item := sliceValue.Index(0)
		if res := reflect.ValueOf(isEqualFunction).Call([]reflect.Value{item, reflect.ValueOf(value)}); res[0].Bool() {
			return 0, nil
		} else {
			return -1, errors.New("item not found")
		}
	} else if len == 2 {
		item0 := sliceValue.Index(0)
		item1 := sliceValue.Index(1)
		if res := reflect.ValueOf(isEqualFunction).Call([]reflect.Value{item0, reflect.ValueOf(value)}); res[0].Bool() {
			return 0, nil
		} else if res := reflect.ValueOf(isEqualFunction).Call([]reflect.Value{item1, reflect.ValueOf(value)}); res[0].Bool() {
			return 1, nil
		} else {
			return -1, errors.New("item not found")
		}
	}

	item := sliceValue.Index(len / 2).Interface()

	var err error = nil
	if err = internal.AreComparable(item, value); err != nil {
		return -1, errors.New("couldn't compare 'value' with all items in passed slice")
	}

	var result int
	if res := reflect.ValueOf(isHigherEqualFunction).Call([]reflect.Value{reflect.ValueOf(item), reflect.ValueOf(value)}); res[0].Bool() {
		if result, err = sortedLastIndexOf(sliceValue.Slice(len/2, len).Interface(), value, isHigherEqualFunction, isEqualFunction); err != nil {
			return -1, err
		}

		return result + (len / 2), nil
	} else {
		if result, err = sortedLastIndexOf(sliceValue.Slice(0, (len/2)+1).Interface(), value, isHigherEqualFunction, isEqualFunction); err != nil {
			return -1, err
		}

		return result, nil
	}
}
