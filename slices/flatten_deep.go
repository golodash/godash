package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Recursively flattens slice.
//
// Complexity: O(n)
//
// n = count of all non-slice type elements of 'slice'
func FlattenDeep(slice interface{}) interface{} {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	sliceItemType := reflect.TypeOf(slice)
	for sliceItemType.Kind() == reflect.Slice {
		sliceItemType = sliceItemType.Elem()
	}

	return recursiveFlattenDeep(slice, sliceItemType).Interface()
}

func recursiveFlattenDeep(slice interface{}, itemType reflect.Type) reflect.Value {
	s := reflect.MakeSlice(reflect.SliceOf(itemType), 0, 0)
	sliceValue := reflect.ValueOf(slice)
	for i := 0; i < sliceValue.Len(); i++ {
		item := reflect.ValueOf(sliceValue.Index(i).Interface())
		if item.Kind() == reflect.Slice {
			s = reflect.AppendSlice(s, recursiveFlattenDeep(item.Interface(), itemType))
		} else {
			s = reflect.Append(s, item)
		}
	}

	return s
}
