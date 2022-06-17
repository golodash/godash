package internal

import (
	"errors"
	"fmt"
	"reflect"
)

var NumberTypes = []reflect.Kind{
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

// Checks if 'slice' interface variable is slice type and
// if 'slice' is nil
func SliceCheck(slice interface{}) error {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return errors.New("not a slice")
	}

	if s.IsNil() {
		return errors.New("slice is nil")
	}

	return nil
}

// Checks if two variables are the same or not
func CheckSameType(var1 interface{}, var2 interface{}) error {
	if !reflect.ValueOf(var1).IsValid() || !reflect.ValueOf(var2).IsValid() {
		return errors.New("invalid values are not allowed")
	}
	if reflect.TypeOf(var1).String() != reflect.TypeOf(var2).String() {
		return errors.New("two variables are not same type")
	}

	return nil
}

// Checks if two given variables are comarable or not
func AreComparable(var1 interface{}, var2 interface{}) error {
	if err := CheckSameType(var1, var2); err != nil {
		return err
	}
	if !reflect.TypeOf(var1).Comparable() {
		return errors.New("two variables are not comparable")
	}

	return nil
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
	v := reflect.ValueOf(input)

	return IsNumberType(v.Kind())
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
func Same(value1 interface{}, value2 interface{}) (condition bool, err error) {
	condition, err = true, nil
	v1 := reflect.ValueOf(value1)
	v2 := reflect.ValueOf(value2)

	// Check for nil and "" and other zero values
	if (!v1.IsValid() && !v2.IsValid()) && (v1.Kind() == v2.Kind()) {
		return true, nil
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
		condition, err = false, nil
		return
	}

	defer func() {
		if r := recover(); r != nil {
			condition, err = false, fmt.Errorf("%s", r)
		}
	}()

	switch v1.Kind() {
	case reflect.Array, reflect.Slice:
		if v1.Len() != v2.Len() {
			condition, err = false, nil
			return
		}
		for i := 0; i < v1.Len(); i++ {
			condition, err = Same(v1.Index(i).Interface(), v2.Index(i).Interface())
			if err != nil || !condition {
				condition, err = false, nil
				return
			}
		}
	case reflect.Map:
		if v1.Len() != v2.Len() {
			condition, err = false, nil
			return
		}

		if len(v1.MapKeys()) != len(v2.MapKeys()) {
			condition, err = false, nil
			return
		}

		keys := v1.MapKeys()
		for i := 0; i < len(v1.MapKeys()); i = i + 1 {
			value1 := v1.MapIndex(keys[i])
			value2 := v2.MapIndex(keys[i])
			if !value1.IsValid() {
				condition, err = false, nil
				return
			}
			if !value2.IsValid() {
				condition, err = false, nil
				return
			}

			if condition, err = Same(value1.Interface(), value2.Interface()); err != nil || !condition {
				condition, err = false, nil
				return
			}
		}
	case reflect.Struct:
		condition, err = value1 == value2, nil
	case reflect.Ptr:
		condition, err = Same(v1.Elem().Interface(), v2.Elem().Interface())
	default:
		if v1.Interface() != v2.Interface() {
			condition, err = false, nil
		}
	}

	return
}
