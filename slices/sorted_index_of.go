package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like IndexOf except that it performs a binary search on a sorted slice.
//
// Complexity: O(log(n))
func SortedIndexOf(slice interface{}, value interface{}) int {
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

	return sortedIndexOf(slice, value, compareLowerEqual, compareIsEqual)
}

// Compare function for SortedIndex function
func compareIsEqual(midValue, value interface{}) bool {
	mid := reflect.ValueOf(midValue)
	v := reflect.ValueOf(value)

	switch v.Kind() {
	case reflect.Float64:
		return v.Interface().(float64) == mid.Interface().(float64)
	case reflect.Float32:
		return v.Interface().(float32) == mid.Interface().(float32)
	case reflect.Int:
		return v.Interface().(int) == mid.Interface().(int)
	case reflect.Int8:
		return v.Interface().(int8) == mid.Interface().(int8)
	case reflect.Int16:
		return v.Interface().(int16) == mid.Interface().(int16)
	case reflect.Int32:
		return v.Interface().(int32) == mid.Interface().(int32)
	case reflect.Int64:
		return v.Interface().(int64) == mid.Interface().(int64)
	case reflect.Uint:
		return v.Interface().(uint) == mid.Interface().(uint)
	case reflect.Uint8:
		return v.Interface().(uint8) == mid.Interface().(uint8)
	case reflect.Uint16:
		return v.Interface().(uint16) == mid.Interface().(uint16)
	case reflect.Uint32:
		return v.Interface().(uint32) == mid.Interface().(uint32)
	case reflect.Uint64:
		return v.Interface().(uint64) == mid.Interface().(uint64)
	case reflect.Uintptr:
		return v.Interface().(uintptr) == mid.Interface().(uintptr)
	}
	return false
}

func sortedIndexOf(slice, value, isLowerEqualFunction, isEqualFunction interface{}) int {
	sliceValue := reflect.ValueOf(slice)
	len := sliceValue.Len()

	if len == 0 {
		return -1
	} else if len == 1 {
		item := reflect.ValueOf(sliceValue.Index(0).Interface())
		if res := reflect.ValueOf(isEqualFunction).Call([]reflect.Value{item, reflect.ValueOf(value)}); res[0].Bool() {
			return 0
		} else {
			return -1
		}
	} else if len == 2 {
		item0 := reflect.ValueOf(sliceValue.Index(0).Interface())
		item1 := reflect.ValueOf(sliceValue.Index(1).Interface())
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
	if res := reflect.ValueOf(isLowerEqualFunction).Call([]reflect.Value{reflect.ValueOf(item), reflect.ValueOf(value)}); res[0].Bool() {
		if result = sortedIndexOf(sliceValue.Slice(0, (len/2)+1).Interface(), value, isLowerEqualFunction, isEqualFunction); result == -1 {
			return -1
		}

		return result
	} else {
		if result = sortedIndexOf(sliceValue.Slice(len/2, len).Interface(), value, isLowerEqualFunction, isEqualFunction); result == -1 {
			return -1
		}

		return result + (len / 2)
	}
}
