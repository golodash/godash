package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like FromPairs except that it accepts two slices, one of property
// identifiers and one of corresponding values.
//
// Complexity: O(n)
func ZipMap(keys, values interface{}) (interface{}, error) {
	if err := internal.SliceCheck(keys); err != nil {
		return nil, err
	}
	if err := internal.SliceCheck(values); err != nil {
		return nil, err
	}

	keysValue := reflect.ValueOf(keys)
	valuesValue := reflect.ValueOf(values)
	if keysValue.Len() != valuesValue.Len() {
		return nil, errors.New("keys and values don't have the same length")
	}
	outputValues := reflect.MakeMapWithSize(reflect.MapOf(keysValue.Type().Elem(), valuesValue.Type().Elem()), keysValue.Len())
	for i := 0; i < keysValue.Len(); i++ {
		outputValues.SetMapIndex(keysValue.Index(i), valuesValue.Index(i))
	}

	return outputValues.Interface(), nil
}
