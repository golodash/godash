package functions

import (
	"reflect"

	"github.com/golodash/godash/internal"
)

// Returns a function that invokes `function`, with passed `inputs` arguments.
//
// Complexity: O(n)
//
// n = length of `inputs`
func WrapFunc(function interface{}, inputs ...interface{}) func() []interface{} {
	if ok := internal.IsFunc(function); !ok {
		panic("`function` input is not a function")
	}

	functionType := reflect.TypeOf(function)
	funcInLen := functionType.NumIn()
	inputsLen := len(inputs)
	defaultPanic := func() { panic("number of `function` inputs don't match with passed inputs") }
	if inputsLen != funcInLen {
		if funcInLen > inputsLen {
			if !functionType.IsVariadic() {
				defaultPanic()
			} else if funcInLen-1 != inputsLen {
				defaultPanic()
			}
		} else {
			if funcInLen != 0 {
				if !functionType.IsVariadic() {
					defaultPanic()
				}
			} else {
				defaultPanic()
			}
		}
	}

	funcInputs := []reflect.Value{}
	for i := 0; i < inputsLen; i++ {
		funcInputs = append(funcInputs, reflect.ValueOf(inputs[i]))
	}

	functionValue := reflect.ValueOf(function)
	return func() []interface{} {
		outputs := functionValue.Call(funcInputs)
		actualOutputs := []interface{}{}
		for i := 0; i < len(outputs); i++ {
			actualOutputs = append(actualOutputs, outputs[i].Interface())
		}

		return actualOutputs
	}
}
