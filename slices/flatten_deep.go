package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Recursively flattens slice.
func FlattenDeep(slice interface{}) ([]interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}

	s, _ := internal.InterfaceToSlice(recursiveFlattenDeep(slice).Interface())

	return s, nil
}

func recursiveFlattenDeep(slice interface{}) reflect.Value {
	s := reflect.ValueOf([]interface{}{})
	sliceValue := reflect.ValueOf(slice)
	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i)
		if val, ok := item.Interface().([]interface{}); ok || item.Kind() == reflect.Slice {
			if val != nil {
				s = reflect.AppendSlice(s, recursiveFlattenDeep(val))
			} else {
				s = reflect.AppendSlice(s, recursiveFlattenDeep(item.Interface()))
			}
		} else {
			s = reflect.Append(s, item)
		}
	}

	return s
}
