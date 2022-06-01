package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

//Returns a slice of unique values, in order, from combine of all given slices.
func Union(slice, slice2 interface{}) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}
	if err := internal.SliceCheck(slice2); err != nil {
		return nil, err
	}
	if reflect.TypeOf(slice).Elem().Kind() != reflect.TypeOf(slice2).Elem().Kind() {
		return nil, errors.New("slices should contain same type of values")
	}

	values := reflect.ValueOf(slice)
	values2 := reflect.ValueOf(slice2)
	for i := 0; i < values2.Len(); i++ {
		values = reflect.Append(values, values2.Index(i))
	}
	union := reflect.MakeSlice(reflect.TypeOf(slice), 0, values.Len())
	m := map[interface{}]bool{}
	for j := 0; j < values.Len(); j++ {
		value := values.Index(j).Interface()
		if _, ok := m[value]; !ok {
			m[value] = true
			union = reflect.Append(union, reflect.ValueOf(value))
		}
	}

	return union.Interface(), nil
}
