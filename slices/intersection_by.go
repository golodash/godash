package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

func IntersectionBy(slice interface{}, function interface{}) ([]interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}

	functionType := reflect.TypeOf(function)
	if functionType.Kind() != reflect.Func {
		return nil, errors.New("`function` has to be function type")
	}
	if functionType.NumIn() != 2 {
		return nil, errors.New("`function` inputs have to have just 2 inputs")
	}
	if functionType.In(0).Kind() != functionType.In(1).Kind() ||
		functionType.In(0).Kind() != reflect.TypeOf(slice).Elem().Kind() &&
			functionType.In(0).Kind() != reflect.Interface {
		return nil, errors.New("`function` inputs have to be the same type as `slice` variable elements or have to be `interface` type")
	}
	if functionType.NumOut() != 1 || functionType.Out(0).Kind() != reflect.Bool {
		return nil, errors.New("`function` function output has to be `bool` type and it has to have just 1 output")
	}

	sliceValue := reflect.ValueOf(slice)
	functionValue := reflect.ValueOf(function)
	values := []interface{}{}
	doubleSeen := []interface{}{}
	for i := 0; i < sliceValue.Len(); i++ {
		item, err := internal.InterfaceToSlice((sliceValue.Index(i).Interface()))
		var ok = false
		if err != nil {
			item, ok = sliceValue.Index(i).Interface().([]interface{})
			if !ok {
				continue
			}
		}

		for j := 0; j < len(item); j++ {
			val := item[j]
			seen := false
			for k := 0; k < len(values); k++ {
				res := functionValue.Call([]reflect.Value{reflect.ValueOf(val), reflect.ValueOf(values[k])})[0]
				if res.Bool() {
					seen = true
					break
				}
			}
			if !seen {
				values = append(values, val)
			} else {
				seen = false
				for k := 0; k < len(doubleSeen); k++ {
					res := functionValue.Call([]reflect.Value{reflect.ValueOf(val), reflect.ValueOf(doubleSeen[k])})[0]
					if res.Bool() {
						seen = true
						break
					}
				}
				if !seen {
					doubleSeen = append(doubleSeen, val)
				}
			}
		}
	}

	return doubleSeen, nil
}
