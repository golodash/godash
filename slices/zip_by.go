package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

func ZipBy(function interface{}, slices ...interface{}) (interface{}, error) {
	var typeOfSlices reflect.Type = nil
	var sizeOfAllSlices int = 0
	var maxLength int = -1
	for i := 0; i < len(slices); i++ {
		if err := internal.SliceCheck(slices[i]); err != nil {
			return nil, err
		}
		if typeOfSlices == nil {
			typeOfSlices = reflect.TypeOf(slices[i])
		}
		if typeOfSlices != reflect.TypeOf(slices[i]) {
			typeOfSlices = reflect.TypeOf([]interface{}{})
		}
		if maxLength == -1 {
			maxLength = reflect.ValueOf(slices[i]).Len()
		}
		if maxLength < reflect.ValueOf(slices[i]).Len() {
			maxLength = reflect.ValueOf(slices[i]).Len()
		}
		sizeOfAllSlices += reflect.ValueOf(slices[i]).Len()
	}
	if typeOfSlices == nil {
		typeOfSlices = reflect.TypeOf([]interface{}{})
	}

	functionValue := reflect.ValueOf(function)
	if functionValue.Kind() != reflect.Func {
		return nil, errors.New("`function` variable has to be `function` type")
	}
	if functionValue.Type().NumOut() != 1 {
		return nil, errors.New("number of `function` output has to be one")
	}
	if functionValue.Type().In(0).Kind() != reflect.Slice && functionValue.Type().In(0).Elem().String() != typeOfSlices.Elem().String() && functionValue.Type().In(0).Kind() != reflect.Interface {
		return nil, errors.New("number of `function` input has to be one and be compatible with inputs slices elements type")
	}

	output := reflect.MakeSlice(reflect.SliceOf(functionValue.Type().Out(0)), 0, sizeOfAllSlices)
	functionInput := []reflect.Value{}
	for j := 0; j < maxLength; j++ {
		for i := 0; i < len(slices); i++ {
			subSlice := reflect.ValueOf(slices[i])
			if j >= subSlice.Len() {
				continue
			}
			functionInput = append(functionInput, subSlice.Index(j))
		}
		output = reflect.Append(output, functionValue.Call(functionInput)[0])
		functionInput = []reflect.Value{}
	}

	return output.Interface(), nil
}
