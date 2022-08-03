package math

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

	outputType := getOutputNumberType(number1, number2)

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
	}

	if internal.CanInt(number1Value.Interface()) {
		var temp int64 = 0
		temp = number1Value.Convert(reflect.TypeOf(temp)).Int() + number2Value.Convert(reflect.TypeOf(temp)).Int()
		output = reflect.ValueOf(temp).Convert(outputType)
	}

	return output.Interface()
}

// Returns a type which has higher rank in between two passed variables.
func getOutputNumberType(number1, number2 interface{}) reflect.Type {
	numbe1Type := reflect.TypeOf(number1)
	numbe2Type := reflect.TypeOf(number2)

	number1Rank := getNumberTypeRank(numbe1Type.Kind())
	number2Rank := getNumberTypeRank(numbe2Type.Kind())

	if number1Rank >= number2Rank {
		return numbe1Type
	} else {
		return numbe2Type
	}
}

// Returns rank of the passed king
func getNumberTypeRank(kind reflect.Kind) int {
	switch kind {
	case reflect.Int8:
		return 0
	case reflect.Uint8:
		return 1
	case reflect.Int16:
		return 2
	case reflect.Uint16:
		return 3
	case reflect.Int32, reflect.Int:
		return 4
	case reflect.Uint32:
		return 5
	case reflect.Int64:
		return 6
	case reflect.Uint64:
		return 7
	case reflect.Uintptr:
		return 8
	case reflect.Float32:
		return 9
	case reflect.Float64:
		return 10
	default:
		return -1
	}
}
