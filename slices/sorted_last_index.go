package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Uses a binary search to determine the Highest index at which value should be
// inserted into slice in order to maintain its sort order.
//
// Complexity: O(log(n))
func SortedLastIndex(slice, value interface{}) int {
	if ok := internal.SliceCheck(slice); !ok {
		panic("passed 'slice' variable is not slice type")
	}

	if res := internal.IsNumber(value); !res {
		panic("'value' is not a number")
	}

	val := reflect.ValueOf(value)
	sType := reflect.TypeOf(slice)
	if !sType.Elem().ConvertibleTo(val.Type()) {
		panic("'value' is not comparable with 'slice'")
	}

	return whereToPutInSliceBiggerEqual(slice, value, compareHigherEqual)
}

// Compare function for SortedLastIndex function
func compareHigherEqual(midValue, value interface{}) bool {
	mid := reflect.ValueOf(midValue)
	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Float64:
		return v.Interface().(float64) >= mid.Interface().(float64)
	case reflect.Float32:
		return v.Interface().(float32) >= mid.Interface().(float32)
	case reflect.Int:
		return v.Interface().(int) >= mid.Interface().(int)
	case reflect.Int8:
		return v.Interface().(int8) >= mid.Interface().(int8)
	case reflect.Int16:
		return v.Interface().(int16) >= mid.Interface().(int16)
	case reflect.Int32:
		return v.Interface().(int32) >= mid.Interface().(int32)
	case reflect.Int64:
		return v.Interface().(int64) >= mid.Interface().(int64)
	case reflect.Uint:
		return v.Interface().(uint) >= mid.Interface().(uint)
	case reflect.Uint8:
		return v.Interface().(uint8) >= mid.Interface().(uint8)
	case reflect.Uint16:
		return v.Interface().(uint16) >= mid.Interface().(uint16)
	case reflect.Uint32:
		return v.Interface().(uint32) >= mid.Interface().(uint32)
	case reflect.Uint64:
		return v.Interface().(uint64) >= mid.Interface().(uint64)
	case reflect.Uintptr:
		return v.Interface().(uintptr) >= mid.Interface().(uintptr)
	}
	return false
}

// Based on binary search, searchs on where to put the sent value in the
// passed slice based on isBiggerEqualFunction result
//
// Complexity: O(log(n))
func whereToPutInSliceBiggerEqual(slice, value, isBiggerEqualFunction interface{}) int {
	sliceValue := reflect.ValueOf(slice)
	len := sliceValue.Len()

	if len == 0 {
		return 0
	} else if len == 1 {
		item := sliceValue.Index(0)
		if res := reflect.ValueOf(isBiggerEqualFunction).Call([]reflect.Value{item, reflect.ValueOf(value)}); res[0].Bool() {
			return 1
		} else {
			return 0
		}
	} else if len == 2 {
		item := sliceValue.Index(1)
		if res := reflect.ValueOf(isBiggerEqualFunction).Call([]reflect.Value{item, reflect.ValueOf(value)}); res[0].Bool() {
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

	if ok := internal.AreComparable(item, value); !ok {
		panic("couldn't compare 'value' with all items in passed slice")
	}

	var result int
	if res := reflect.ValueOf(isBiggerEqualFunction).Call([]reflect.Value{reflect.ValueOf(item), reflect.ValueOf(value)}); res[0].Bool() {
		if result = whereToPutInSliceBiggerEqual(sliceValue.Slice(len/2, len).Interface(), value, isBiggerEqualFunction); result == -1 {
			return -1
		}

		return result + (len / 2)
	} else {
		if result = whereToPutInSliceBiggerEqual(sliceValue.Slice(0, (len/2)+1).Interface(), value, isBiggerEqualFunction); result == -1 {
			return -1
		}

		return result
	}
}
