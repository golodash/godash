package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like `difference` except that it accepts a custom
// comparator function which is invoked to compare elements of slices to values.
//
// comparator has to indicate if two given variables as inputs are equal or not.
//
// comparator example function:
//
// 	func compareDifferenceByTest(value1, value2 interface{}) bool {
// 		v1 := reflect.ValueOf(value1).Int()
// 		v2 := reflect.ValueOf(value2).Int()
// 		return v1 == v2
// 	}
func DifferenceBy(slice interface{}, notIncluded interface{}, comparator interface{}) ([]interface{}, error) {
	if err := internal.CheckSameType(slice, notIncluded); err != nil {
		return nil, err
	}

	if err1, err2 := internal.SliceCheck(slice), internal.SliceCheck(notIncluded); err1 != nil || err2 != nil {
		if err2 != nil {
			return nil, err2
		}
		return nil, err1
	}

	comparatorType := reflect.TypeOf(comparator)
	if comparatorType.Kind() != reflect.Func {
		return nil, errors.New("`comparator` has to be function type")
	}
	if comparatorType.NumIn() != 2 {
		return nil, errors.New("`comparator` function inputs has to have just 2 inputs")
	}
	if comparatorType.In(0).Kind() != comparatorType.In(1).Kind() ||
		comparatorType.In(0).Kind() != reflect.TypeOf(slice).Elem().Kind() &&
			comparatorType.In(0).Kind() != reflect.Interface {
		return nil, errors.New("`comparator` function inputs have to be the same type as `slice` and `notIncluded` variables or have to be `interface` type")
	}
	if comparatorType.NumOut() != 1 || comparatorType.Out(0).Kind() != reflect.Bool {
		return nil, errors.New("`comparator` function output has to be `bool` type and it has to have just 1 output")
	}

	s, err := internal.InterfaceToSlice(slice)
	if err != nil {
		return nil, err
	}

	functionValue := reflect.ValueOf(comparator)
	notIn := reflect.ValueOf(notIncluded)

	for i := len(s) - 1; i > -1; i-- {
		if i >= len(s) {
			continue
		}
	firstLoop:
		for j := 0; j < notIn.Len(); j++ {
			res := functionValue.Call([]reflect.Value{reflect.ValueOf(s[i]), notIn.Index(j)})

			if res[0].Bool() {
				if i != 0 && i+1 < len(s) {
					s = append(s[0:i], s[i+1:]...)
				} else if i == 0 {
					s = s[i+1:]
				} else if i+1 >= len(s) {
					s = s[0:i]
				}
				i++
				break firstLoop
			}
		}
	}

	return s, nil
}
