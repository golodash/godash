package slices

import (
	"fmt"
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method returns an object composed from key-value pairs.
//
// Complexity: O(n)
func FromPairs(slice interface{}) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
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
		if err := internal.SliceCheck(item.Interface()); err != nil {
			return nil, fmt.Errorf("item in index %d is not even a slice", i)
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

	return output.Interface(), nil
}
