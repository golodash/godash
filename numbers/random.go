package maths

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/golodash/godash/internal"
)

// Produces a random number between the inclusive lower and upper bounds.If floating is true, or either lower or upper are floats, a floating-point number is returned instead of an integer.
//
// Complexity: O(1)
func Random(lower, upper interface{}, floating bool) interface{} {

	if !internal.IsNumber(lower) {
		panic("'number' is not a number")
	}
	if !internal.IsNumber(upper) {
		panic("'lower' is not a number")
	}

	rand.Seed(time.Now().UnixNano())

	if floating == true {
		lowerValue := reflect.ValueOf(lower).Convert(reflect.TypeOf(1.0)).Float()
		upperValue := reflect.ValueOf(upper).Convert(reflect.TypeOf(1.0)).Float()
		if lowerValue <= upperValue {
			return lowerValue + rand.Float64()*(upperValue-lowerValue)
		} else {
			return upperValue + rand.Float64()*(lowerValue-upperValue)
		}
	} else {
		lowerValue := reflect.ValueOf(lower).Int()
		upperValue := reflect.ValueOf(upper).Int()
		if lowerValue <= upperValue {
			return rand.Intn(upperValue-lowerValue) + lowerValue
		} else {
			return rand.Intn(upperValue-lowerValue) + lowerValue
		}
	}
}
