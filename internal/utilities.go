package internal

import (
	"fmt"
	"reflect"
	"testing"
)

// Used to inform difference of two variable, which one is
// bigger, smaller or maybe they have the same value
type Diff uint

const (
	Lower Diff = iota
	Higher
	Equal
)

var (
	NumberTypes = []reflect.Kind{
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.Uintptr,
	}
)

// Checks if 'slice' interface variable is slice type and if 'slice' is nil
func SliceCheck(slice interface{}) bool {
	s := reflect.ValueOf(slice)
	return s.IsValid() && s.Kind() == reflect.Slice && !s.IsNil()
}

// Checks if two variables are the same or not
func CheckSameType(var1 interface{}, var2 interface{}) bool {
	return (reflect.ValueOf(var1).IsValid() && reflect.ValueOf(var2).IsValid()) && reflect.TypeOf(var1).String() == reflect.TypeOf(var2).String()
}

// Checks if two given variables are comarable or not
func AreComparable(var1 interface{}, var2 interface{}) bool {
	return reflect.TypeOf(var1).Comparable() && CheckSameType(var1, var2)
}

// Returns a unique list of integers
func UniqueInt(s []int) []int {
	inResult := make(map[int]bool)
	var result []int
	for _, num := range s {
		if _, ok := inResult[num]; !ok {
			inResult[num] = true
			result = append(result, num)
		}
	}
	return result
}

// Returns true if passed variable is a number
func IsNumber(input interface{}) bool {
	return IsNumberType(reflect.ValueOf(input).Kind())
}

// Returns true if passed kind is a number
func IsNumberType(input reflect.Kind) bool {
	for _, value := range NumberTypes {
		if value == input {
			return true
		}
	}

	return false
}

// Determines if passed variables are exactly the same
func Same(value1 interface{}, value2 interface{}) (condition bool) {
	condition = true
	v1 := reflect.ValueOf(value1)
	v2 := reflect.ValueOf(value2)

	// Check for nil and "" and other zero values
	if (!v1.IsValid() && !v2.IsValid()) && (v1.Kind() == v2.Kind()) {
		return
	}

	if v1.Kind() != v2.Kind() {
		if v1.Kind() == reflect.Interface && v2.Kind() != reflect.Interface {
			if v1.CanConvert(v2.Type()) {
				v1 = v1.Convert(v2.Type())
			}
		} else if v2.Kind() == reflect.Interface && v1.Kind() != reflect.Interface {
			if v2.CanConvert(v1.Type()) {
				v2 = v2.Convert(v1.Type())
			}
		}
		if v1.Kind() == v2.Kind() && v1.Interface() == v2.Interface() {
			return
		}
		condition = false
		return
	}

	defer func() {
		if r := recover(); r != nil {
			condition = false
		}
	}()

	switch v1.Kind() {
	case reflect.Array, reflect.Slice:
		if v1.Len() != v2.Len() {
			condition = false
			return
		}
		for i := 0; i < v1.Len(); i++ {
			condition = Same(v1.Index(i).Interface(), v2.Index(i).Interface())
			if !condition {
				condition = false
				return
			}
		}
	case reflect.Map:
		if v1.Len() != v2.Len() {
			condition = false
			return
		}

		if len(v1.MapKeys()) != len(v2.MapKeys()) {
			condition = false
			return
		}

		keys := v1.MapKeys()
		for i := 0; i < len(v1.MapKeys()); i = i + 1 {
			value1 := v1.MapIndex(keys[i])
			value2 := v2.MapIndex(keys[i])
			if !value1.IsValid() {
				condition = false
				return
			}
			if !value2.IsValid() {
				condition = false
				return
			}

			if condition = Same(value1.Interface(), value2.Interface()); !condition {
				condition = false
				return
			}
		}
	case reflect.Struct:
		condition = value1 == value2
	case reflect.Ptr:
		condition = Same(v1.Elem().Interface(), v2.Elem().Interface())
	default:
		if v1.Interface() != v2.Interface() {
			condition = false
		}
	}

	return
}

// Used in test cases to prevent code breaking
func DeferTestCases(t *testing.T, expected interface{}) {
	err := recover()

	if err != nil && expected != nil {
		t.Errorf("wanted = %v, err = %s", expected, err)
	}
}

// Returns true if passed argument type is int
func CanInt(number interface{}) bool {
	switch reflect.TypeOf(number).Kind() {
	case reflect.Int:
		return true
	case reflect.Int8:
		return true
	case reflect.Int16:
		return true
	case reflect.Int32:
		return true
	case reflect.Int64:
		return true
	}
	return false
}

// Returns true if passed argument type is uint
func CanUint(number interface{}) bool {
	switch reflect.TypeOf(number).Kind() {
	case reflect.Uint:
		return true
	case reflect.Uint8:
		return true
	case reflect.Uint16:
		return true
	case reflect.Uint32:
		return true
	case reflect.Uint64:
		return true
	case reflect.Uintptr:
		return true
	}
	return false
}

// Returns true if passed argument type is float
func CanFloat(number interface{}) bool {
	switch reflect.TypeOf(number).Kind() {
	case reflect.Float32:
		return true
	case reflect.Float64:
		return true
	}
	return false
}

// Returns a type which has higher rank in between two passed number variables.
func GetOutputNumberType(number1, number2 interface{}) reflect.Type {
	numbe1Type := reflect.TypeOf(number1)
	numbe2Type := reflect.TypeOf(number2)

	number1Rank := GetNumberTypeRank(numbe1Type.Kind())
	number2Rank := GetNumberTypeRank(numbe2Type.Kind())

	if number1Rank >= number2Rank {
		return numbe1Type
	} else {
		return numbe2Type
	}
}

// Returns rank of the passed kind
func GetNumberTypeRank(kind reflect.Kind) int {
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

func CompareNumbers(number1, number2 interface{}) Diff {
	number1Value := reflect.ValueOf(reflect.ValueOf(number1).Interface())
	number2Value := reflect.ValueOf(reflect.ValueOf(number2).Interface())
	number1 = number1Value.Interface()
	number2 = number2Value.Interface()

	if IsNumber(number1) && IsNumber(number2) {
		if CanFloat(number1) || CanFloat(number2) {
			floatType := reflect.TypeOf(1.0)
			n1 := number1Value.Convert(floatType).Float()
			n2 := number2Value.Convert(floatType).Float()
			if n1 > n2 {
				return Higher
			} else if n1 < n2 {
				return Lower
			} else {
				return Equal
			}
		} else {
			intType := reflect.TypeOf(1)
			n1 := number1Value.Convert(intType).Int()
			n2 := number2Value.Convert(intType).Int()
			if n1 > n2 {
				return Higher
			} else if n1 < n2 {
				return Lower
			} else {
				return Equal
			}
		}
	} else {
		panic(fmt.Sprintf("%s, %s are not numbers to compare", number1, number2))
	}
}

func CustomIsSeparator(letter rune, separators []rune) bool {
	for i := 0; i < len(separators); i++ {
		if letter == separators[i] {
			return true
		}
	}

	return false
}
