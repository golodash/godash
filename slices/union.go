package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Returns a slice of unique values, in order, from combine of all given slices.
//
// Complexity: O(n)
//
// n = length of both slices combined
func Union(slice1, slice2 interface{}) interface{} {
	if !internal.SliceCheck(slice1) {
		panic("passed 'slice1' variable is not slice type")
	}
	if !internal.SliceCheck(slice2) {
		panic("passed 'slice2' variable is not slice type")
	}

	sameType := false
	values := reflect.ValueOf(slice1)
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
		value := values.Index(i).Interface()
		if _, ok := m[value]; !ok {
			m[value] = true
			union = reflect.Append(union, reflect.ValueOf(value))
		}
	}
	for i := 0; i < values2.Len(); i++ {
		value := values2.Index(i).Interface()
		if _, ok := m[value]; !ok {
			m[value] = true
			union = reflect.Append(union, reflect.ValueOf(value))
		}
	}

	return union.Interface()
}
