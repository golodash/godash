package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Returns a slice of unique values, in order, from combine of all given slices.
//
// example for 'function':
//
//  func makeInt(input interface{}) interface{} {
//    return int(input.(float64))
//  }
//
// Complexity: O(n)
//
// n = length of both slices combined
func UnionBy(slice, slice2 interface{}, function func(interface{}) interface{}) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}
	if err := internal.SliceCheck(slice2); err != nil {
		return nil, err
	}

	sameType := false
	values := reflect.ValueOf(slice)
	values2 := reflect.ValueOf(slice2)
	if values.Type().String() == values2.Type().String() {
		sameType = true
	}

	var union reflect.Value = reflect.Value{}
	if sameType {
		union = reflect.MakeSlice(values.Type(), 0, values.Len()+values2.Len())
	} else {
		union = reflect.MakeSlice(reflect.TypeOf([]interface{}{}), 0, values.Len()+values2.Len())
	}

	m := map[interface{}]bool{}
	for i := 0; i < values.Len(); i++ {
		value := function(values.Index(i).Interface())
		if _, ok := m[value]; !ok {
			m[value] = true
			union = reflect.Append(union, values.Index(i))
		}
	}
	for i := 0; i < values2.Len(); i++ {
		value := function(values2.Index(i).Interface())
		if _, ok := m[value]; !ok {
			m[value] = true
			union = reflect.Append(union, values2.Index(i))
		}
	}

	return union.Interface(), nil
}
