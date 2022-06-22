package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like FromPairs except that it accepts two slices, one of property
// identifiers and one of corresponding values.
//
// Complexity: O(n)
func ZipMap(keys, values interface{}) interface{} {
	if ok := internal.SliceCheck(keys); !ok {
		panic("passed 'keys' variable is not slice type")
	}
	if ok := internal.SliceCheck(values); !ok {
		panic("passed 'values' variable is not slice type")
	}

	keysValue := reflect.ValueOf(keys)
	valuesValue := reflect.ValueOf(values)
	if keysValue.Len() != valuesValue.Len() {
		panic("keys and values don't have the same length")
	}
	outputValues := reflect.MakeMapWithSize(reflect.MapOf(keysValue.Type().Elem(), valuesValue.Type().Elem()), keysValue.Len())
	for i := 0; i < keysValue.Len(); i++ {
		outputValues.SetMapIndex(keysValue.Index(i), valuesValue.Index(i))
	}

	return outputValues.Interface()
}
