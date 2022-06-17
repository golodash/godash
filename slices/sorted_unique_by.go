package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like 'UniqueBy' except that it's designed and optimized
// for sorted slices.
//
// It accepts a function which is invoked for each element in slice to
// generate the criterion by which uniqueness is computed.
//
// example for 'function':
//
//  func makeInt(input interface{}) interface{} {
//    return int(input.(float64))
//  }
//
// Complexity: O(n)
func SortedUniqueBy(slice interface{}, function func(interface{}) interface{}) (interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}

	sliceType := reflect.TypeOf(slice)
	sliceValue := reflect.ValueOf(slice)
	output := reflect.MakeSlice(sliceType, 0, sliceValue.Len())

	sliceItemType := sliceType.Elem()
	tempMap := reflect.MakeMap(reflect.MapOf(sliceItemType, reflect.TypeOf(true)))
	for i := 0; i < sliceValue.Len(); i++ {
		item := reflect.ValueOf(sliceValue.Index(i).Interface())
		key := function(item.Interface())
		if exist := tempMap.MapIndex(reflect.ValueOf(key)); !exist.IsValid() {
			tempMap.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(true))
			output = reflect.Append(output, item)
		}
	}

	return output.Interface(), nil
}
