package slices

import (
	"fmt"
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method returns an object composed from key-value pairs.
//
// Complexity: O(n)
func FromPairs(slice interface{}) interface{} {
	if ok := internal.SliceCheck(slice); !ok {
		panic("passed 'slice' variable is not slice type")
	}

	sliceItemType := reflect.TypeOf(slice)
	if sliceItemType = sliceItemType.Elem(); sliceItemType.Kind() == reflect.Slice {
		if sliceItemType = sliceItemType.Elem(); sliceItemType.Kind() == reflect.Slice {
			_ = 0
		}
	}

	sliceValue := reflect.ValueOf(slice)
	output := reflect.MakeMap(reflect.MapOf(sliceItemType, sliceItemType))
	for i := 0; i < sliceValue.Len(); i++ {
		item := reflect.ValueOf(sliceValue.Index(i).Interface())
		if ok := internal.SliceCheck(item.Interface()); !ok {
			panic(fmt.Sprintf("item in index %d is not even a slice", i))
		}

		if item.Len() == 2 {
			if key, ok := item.Index(0).Interface().(string); ok {
				output.SetMapIndex(reflect.ValueOf(key), item.Index(1))
			}
		} else if item.Len() == 1 {
			if key, ok := item.Index(0).Interface().(string); ok {
				output.SetMapIndex(reflect.ValueOf(key), reflect.Zero(sliceItemType))
			}
		}
	}

	return output.Interface()
}
