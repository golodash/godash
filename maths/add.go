package maths

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Adds two numbers.
//
// Complexity: O(1)
func Add(number1, number2 interface{}) interface{} {
	if !internal.IsNumber(number1) {
		panic("'number1' is not a number")
	}
	if !internal.IsNumber(number2) {
		panic("'number2' is not a number")
	}

	outputType := internal.GetOutputNumberType(number1, number2)

	number1Value := reflect.ValueOf(number1)
	number2Value := reflect.ValueOf(number2)

	if number1Value.Kind() != outputType.Kind() {
		number1Value = number1Value.Convert(outputType)
	} else if number2Value.Kind() != outputType.Kind() {
		number2Value = number2Value.Convert(outputType)
	}

	output := reflect.Zero(outputType)
	if internal.CanFloat(number1Value.Interface()) {
		var temp float64 = 0
		temp = number1Value.Convert(reflect.TypeOf(temp)).Float() + number2Value.Convert(reflect.TypeOf(temp)).Float()
		output = reflect.ValueOf(temp).Convert(outputType)
	} else if internal.CanUint(number1Value.Interface()) {
		var temp uint64 = 0
		temp = number1Value.Convert(reflect.TypeOf(temp)).Uint() + number2Value.Convert(reflect.TypeOf(temp)).Uint()
		output = reflect.ValueOf(temp).Convert(outputType)
	} else if internal.CanInt(number1Value.Interface()) {
		var temp int64 = 0
		temp = number1Value.Convert(reflect.TypeOf(temp)).Int() + number2Value.Convert(reflect.TypeOf(temp)).Int()
		output = reflect.ValueOf(temp).Convert(outputType)
	}

	return output.Interface()
}
