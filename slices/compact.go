package slices

import (
	"reflect"

	"github.com/golodash/godash/generals"
	"github.com/golodash/godash/internal"
)

// Removes falsey items from slice except values you mentioned.
//
// Falsey items are {"", nil, 0, false}
//
// Complexity: O(n)
func Compact(slice, excepts interface{}) interface{} {
	if !internal.SliceCheck(slice) {
		panic("passed 'slice' variable is not slice type")
	}

	exceptsValue := reflect.ValueOf(excepts)
	if exceptsValue.Kind() != reflect.Slice && excepts != nil {
		panic("just slice accepted as 'excepts' value")
	}
	if !exceptsValue.IsValid() {
		exceptsValue = reflect.MakeSlice(reflect.TypeOf([]interface{}{}), 0, 0)
	}

	defaultFalsey := []interface{}{"", nil, 0, false}
	falsey := []interface{}{}
	for i := 0; i < len(defaultFalsey); i++ {
		remain := true
		for j := 0; j < exceptsValue.Len(); j++ {
			if defaultFalsey[i] == exceptsValue.Index(j).Interface() {
				remain = false
			}
		}
		if remain {
			falsey = append(falsey, defaultFalsey[i])
		}
	}

	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	result := reflect.MakeSlice(reflect.TypeOf(slice), 0, length)
	j := 0
	for i := 0; i < length; i++ {
		for k := 0; k < len(falsey); k++ {
			if generals.Same(sliceValue.Index(i).Interface(), falsey[k]) {
				if i == j {
					j = i + 1
					continue
				}
				result = reflect.AppendSlice(result, sliceValue.Slice(j, i))
				j = i + 1
			}
		}
	}

	if j < sliceValue.Len() {
		result = reflect.AppendSlice(result, sliceValue.Slice(j, length))
	}

	return result.Interface()
}
