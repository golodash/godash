package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Creates a slice of unique values that are included in all given slices
// for equality comparisons. The order and references of result values are
// determined by the first slice.
func Intersection(slices interface{}) (interface{}, error) {
	if err := internal.SliceCheck(slices); err != nil {
		return nil, err
	}

	sliceItemType := reflect.TypeOf(slices)
	if sliceItemType = sliceItemType.Elem(); sliceItemType.Kind() == reflect.Slice {
		sliceItemType = sliceItemType.Elem()
	}

	sliceValue := reflect.ValueOf(slices)
	seenMap := reflect.MakeMap(reflect.MapOf(sliceItemType, reflect.TypeOf(false)))
	outputSlice := reflect.MakeSlice(reflect.SliceOf(sliceItemType), 0, sliceValue.Len())
	for i := 0; i < sliceValue.Len(); i++ {
		subSlice := reflect.ValueOf(sliceValue.Index(i).Interface())
		if err := internal.SliceCheck(subSlice.Interface()); err != nil {
			continue
		}

		for j := 0; j < subSlice.Len(); j++ {
			item := reflect.ValueOf(subSlice.Index(j).Interface())
			var value reflect.Value = reflect.Value{}
			if value = seenMap.MapIndex(item); value.IsValid() && !value.IsZero() {
				continue
			}

			outputSlice = reflect.Append(outputSlice, item)
			seenMap.SetMapIndex(item, reflect.ValueOf(true))
		}
	}

	return outputSlice.Interface(), nil
}
