package slices

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// This method is like Zip except that it accepts a function to specify
// how grouped values should be combined.
//
// example for 'function':
//
//  func makeInt(input interface{}) interface{} {
//    return int(input.(float64))
//  }
//
// Complexity: O(n)
func ZipBy(slices interface{}, function func(interface{}) interface{}) (interface{}, error) {
	if err := internal.SliceCheck(slices); err != nil {
		return nil, err
	}

	// Calculating output type
	var output = reflect.Value{}
	sliceItemType := reflect.TypeOf(slices)
	slicesValue := reflect.ValueOf(slices)
	for sliceItemType.Kind() == reflect.Slice {
		sliceItemType = sliceItemType.Elem()
	}
	if sliceItemType.Kind() != reflect.Interface {
		output = reflect.MakeSlice(reflect.TypeOf(slices).Elem(), 0, slicesValue.Len()/2)
	} else {
		output = reflect.MakeSlice(reflect.TypeOf([][]interface{}{}), 0, slicesValue.Len()/2)
	}

	// Calculating length
	sizeOfAllSlices := 0
	maxLength := -1
	for i := 0; i < slicesValue.Len(); i++ {
		item := slicesValue.Index(i).Interface()
		itemValue := reflect.ValueOf(item)
		if err := internal.SliceCheck(item); err != nil {
			return nil, err
		}
		if maxLength < itemValue.Len() {
			maxLength = itemValue.Len()
		}
		sizeOfAllSlices += itemValue.Len()
	}

	functionInput := reflect.MakeSlice(reflect.TypeOf(slices).Elem(), 0, 0)
	for j := 0; j < maxLength; j++ {
		for i := 0; i < slicesValue.Len(); i++ {
			subSlice := reflect.ValueOf(slicesValue.Index(i).Interface())
			if j >= subSlice.Len() {
				continue
			}
			functionInput = reflect.Append(functionInput, subSlice.Index(j))
		}
		funcOutput := reflect.ValueOf(function(functionInput.Interface()))

		output = reflect.Append(output, funcOutput)
		functionInput = reflect.Zero(functionInput.Type())
	}

	return output.Interface(), nil
}
