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

	return getTypeRecursiveFlattenDeep(slice), nil
}

func getTypeRecursiveFlattenDeep(slice interface{}) interface{} {
	sliceType := reflect.TypeOf(slice)
	for sliceType.Kind() == reflect.Slice {
		sliceType = sliceType.Elem()
	}

	outputItems := recursiveFlattenDeep(slice).Interface().([]interface{})
	if sliceType.Kind() != reflect.Interface {
		output := reflect.MakeSlice(reflect.SliceOf(sliceType), 0, len(outputItems))
		for i := 0; i < len(outputItems); i++ {
			output = reflect.Append(output, reflect.ValueOf(outputItems[i]))
		}

		return output.Interface()
	} else {
		return outputItems
	}
}

func recursiveFlattenDeep(slice interface{}) reflect.Value {
	s := reflect.ValueOf([]interface{}{})
	sliceValue := reflect.ValueOf(slice)
	for i := 0; i < sliceValue.Len(); i++ {
		item := reflect.ValueOf(sliceValue.Index(i).Interface())
		if item.Kind() == reflect.Slice {
			s = reflect.AppendSlice(s, recursiveFlattenDeep(item.Interface()))
		} else {
			s = reflect.Append(s, item)
		}
	}

	return s
}
