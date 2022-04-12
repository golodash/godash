package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Flattens slice a single level deep.
func FlattenDepth(slice interface{}, depth int) ([]interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}

	if depth == 0 {
		return internal.InterfaceToSlice(slice)
	} else if depth < 0 {
		s, err := internal.InterfaceToSlice(slice)
		if err != nil {
			return nil, err
		}
		return negativeFlatten(s, -depth), nil
	} else {
		return internal.InterfaceToSlice(recursiveFlattenDepth(reflect.ValueOf(slice), depth).Interface())
	}
}

func negativeFlatten(slice []interface{}, depth int) []interface{} {
	if depth > 0 {
		slice = []interface{}{negativeFlatten(slice, depth-1)}
	}
	return slice
}

func recursiveFlattenDepth(slice reflect.Value, depth int) reflect.Value {
	s := reflect.ValueOf([]interface{}{})
	for i := 0; i < slice.Len(); i++ {
		item := slice.Index(i)
		if val, ok := item.Interface().([]interface{}); ok || item.Kind() == reflect.Slice {
			if val != nil {
				if depth != 0 {
					s = reflect.AppendSlice(s, recursiveFlattenDepth(reflect.ValueOf(val), depth-1))
				} else {
					s = reflect.Append(s, reflect.ValueOf(val))
				}
			} else {
				if depth != 0 {
					s = reflect.AppendSlice(s, recursiveFlattenDepth(item, depth-1))
				} else {
					s = reflect.Append(s, item)
				}
			}
		} else {
			s = reflect.Append(s, item)
		}
	}

	return s
}
