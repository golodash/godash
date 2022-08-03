package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Uses a binary search to determine the lowest index at which value should be
// inserted into slice in order to maintain its sort order.
//
// Complexity: O(log(n))
func SortedIndex(slice, value interface{}) int {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	if !internal.IsNumber(value) {
		panic("'value' is not a number")
	}

	val := reflect.ValueOf(value)
	sType := reflect.TypeOf(slice)
	if !sType.Elem().ConvertibleTo(val.Type()) {
		panic("'value' is not comparable with 'slice'")
	}

	if reflect.ValueOf(slice).Len() == 0 {
		return 0
	}

	return whereToPutInSliceLowerEqual(slice, value, compareLowerEqual)
}

// Compare function for SortedIndex function
func compareLowerEqual(midValue, value interface{}) bool {
	mid := reflect.ValueOf(midValue)
	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Float64:
		return v.Interface().(float64) <= mid.Interface().(float64)
	case reflect.Float32:
		return v.Interface().(float32) <= mid.Interface().(float32)
	case reflect.Int:
		return v.Interface().(int) <= mid.Interface().(int)
	case reflect.Int8:
		return v.Interface().(int8) <= mid.Interface().(int8)
	case reflect.Int16:
		return v.Interface().(int16) <= mid.Interface().(int16)
	case reflect.Int32:
		return v.Interface().(int32) <= mid.Interface().(int32)
	case reflect.Int64:
		return v.Interface().(int64) <= mid.Interface().(int64)
	case reflect.Uint:
		return v.Interface().(uint) <= mid.Interface().(uint)
	case reflect.Uint8:
		return v.Interface().(uint8) <= mid.Interface().(uint8)
	case reflect.Uint16:
		return v.Interface().(uint16) <= mid.Interface().(uint16)
	case reflect.Uint32:
		return v.Interface().(uint32) <= mid.Interface().(uint32)
	case reflect.Uint64:
		return v.Interface().(uint64) <= mid.Interface().(uint64)
	case reflect.Uintptr:
		return v.Interface().(uintptr) <= mid.Interface().(uintptr)
	}
	return false
}

// Based on binary search, searches on where to put the sent value in the passed
// slice based on isLowerEqualFunction result.
//
// Complexity: O(log(n))
func whereToPutInSliceLowerEqual(slice, value, isLowerEqualFunction interface{}) int {
	sliceValue := reflect.ValueOf(slice)
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
	if res := reflect.ValueOf(isLowerEqualFunction).Call([]reflect.Value{reflect.ValueOf(item), reflect.ValueOf(value)}); res[0].Bool() {
		if result = whereToPutInSliceLowerEqual(sliceValue.Slice(0, (len/2)+1).Interface(), value, isLowerEqualFunction); result == -1 {
			return -1
		}

		return result
	} else {
		if result = whereToPutInSliceLowerEqual(sliceValue.Slice(len/2, len).Interface(), value, isLowerEqualFunction); result == -1 {
			return -1
		}

		return result + (len / 2)
	}
}
