package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// This function creates a duplicate-free version of an slice.
//
// This method is designed and optimized for sorted slices.
//
// Complexity: O(n)
func SortedUnique(slice interface{}) interface{} {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	sliceItemType := reflect.TypeOf(slice).Elem()
	sliceValue := reflect.ValueOf(slice)
	tempMap := reflect.MakeMap(reflect.MapOf(sliceItemType, reflect.TypeOf(true)))
	output := reflect.MakeSlice(reflect.TypeOf(slice), 0, sliceValue.Len())

	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i)
		if exist := tempMap.MapIndex(item); !exist.IsValid() {
			tempMap.SetMapIndex(item, reflect.ValueOf(true))
			output = reflect.Append(output, item)
		}
	}

	return output.Interface()
}
