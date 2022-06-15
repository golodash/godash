package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Recursively flattens slice.
func FlattenDeep(slice interface{}) (interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}

	sliceItemType := reflect.TypeOf(slice)
	for sliceItemType.Kind() == reflect.Slice {
		sliceItemType = sliceItemType.Elem()
	}

	return recursiveFlattenDeep(slice, sliceItemType).Interface(), nil
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
