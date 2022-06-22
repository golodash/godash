package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like LastIndexOf except that it performs a
// binary search on a sorted slice.
//
// Complexity: O(log(n))
func SortedLastIndexOf(slice, value interface{}) int {
	if ok := internal.SliceCheck(slice); !ok {
		panic("passed 'slice' variable is not slice type")
	}

	sType := reflect.TypeOf(slice)
	val := reflect.ValueOf(value)
	if val.Type().Kind() != sType.Elem().Kind() && sType.Elem().Kind() != reflect.Interface {
		panic("'value' is not compatible with 'slice' elements")
	}

	return sortedLastIndexOf(slice, value, compareHigherEqual, compareIsEqual)
}

// Complexity: O(log(n))
func sortedLastIndexOf(slice, value, isHigherEqualFunction, isEqualFunction interface{}) int {
	sliceValue := reflect.ValueOf(slice)
	len := sliceValue.Len()

	if len == 0 {
		return -1
	} else if len == 1 {
		item := sliceValue.Index(0)
		if res := reflect.ValueOf(isEqualFunction).Call([]reflect.Value{item, reflect.ValueOf(value)}); res[0].Bool() {
			return 0
		} else {
			return -1
		}
	} else if len == 2 {
		item0 := sliceValue.Index(0)
		item1 := sliceValue.Index(1)
		if res := reflect.ValueOf(isEqualFunction).Call([]reflect.Value{item0, reflect.ValueOf(value)}); res[0].Bool() {
			return 0
		} else if res := reflect.ValueOf(isEqualFunction).Call([]reflect.Value{item1, reflect.ValueOf(value)}); res[0].Bool() {
			return 1
		} else {
			return -1
		}
	}

	item := sliceValue.Index(len / 2).Interface()

	if ok := internal.AreComparable(item, value); !ok {
		panic("couldn't compare 'value' with all items in passed slice")
	}

	var result int
	if res := reflect.ValueOf(isHigherEqualFunction).Call([]reflect.Value{reflect.ValueOf(item), reflect.ValueOf(value)}); res[0].Bool() {
		if result = sortedLastIndexOf(sliceValue.Slice(len/2, len).Interface(), value, isHigherEqualFunction, isEqualFunction); result == -1 {
			return -1
		}

		return result + (len / 2)
	} else {
		if result = sortedLastIndexOf(sliceValue.Slice(0, (len/2)+1).Interface(), value, isHigherEqualFunction, isEqualFunction); result == -1 {
			return -1
		}

		return result
	}
}
