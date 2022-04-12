package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Flattens slice a single level deep.
func Flatten(slice interface{}) ([]interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}

	s := []interface{}{}
	sliceValue := reflect.ValueOf(slice)
	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i)
		if val, ok := item.Interface().([]interface{}); ok || item.Kind() == reflect.Slice {
			if val != nil {
				for j := 0; j < len(val); j++ {
					s = append(s, val[j])
				}
			} else {
				for j := 0; j < item.Len(); j++ {
					s = append(s, item.Index(j).Interface())
				}
			}
		} else {
			s = append(s, item.Interface())
		}
	}

	return s, nil
}
