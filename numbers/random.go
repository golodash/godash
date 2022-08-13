package maths

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/golodash/godash/internal"
)

// Produces a random number between the inclusive 'lower' and exclusive 'upper' bounds.
// If floating is true, a floating-point number is returned instead of an integer.
//
// If floating is true, a float64 number type is returned.
// If floating is false, an int64 number type is returned.
//
// Complexity: O(1)
func Random(lower, upper interface{}, floating bool) interface{} {
	if !internal.IsNumber(lower) {
		panic("'lower' is not a number")
	}
	if !internal.IsNumber(upper) {
		panic("'upper' is not a number")
	}

	rand.Seed(time.Now().UnixNano())

	lowerFloat := reflect.ValueOf(lower).Convert(reflect.TypeOf(1.0)).Float()
	upperFloat := reflect.ValueOf(upper).Convert(reflect.TypeOf(1.0)).Float()

	if lowerFloat > upperFloat {
		panic("'upper' has to be higher or equal to 'lower'")
	}

	output := lowerFloat + rand.Float64()*(upperFloat-lowerFloat)

	if !floating {
		return int64(output)
	}
	return output
}
