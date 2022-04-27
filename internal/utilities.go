package internal

import (
	"errors"
	"reflect"
	"runtime"
	"strings"
)

// Send a function as input param to this function and
// get the package name of that function as string
func GetPackageName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	strs = strings.Split(strs[len(strs)-2], "/")
	return strs[len(strs)-1]
}

// Send a function as input param to this function and
// get the function name as string
func GetFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

// Converts interface{} to []interface{}
func InterfaceToSlice(slice interface{}) ([]interface{}, error) {
	err := SliceCheck(slice)
	if err != nil {
		return nil, err
	}

	s := reflect.ValueOf(slice)
	result := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		result[i] = s.Index(i).Interface()
	}

	return result, nil
}

// Checks if `slice` interface variable is slice type and
// if `slice` is nil
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

func IsNumber(input interface{}) bool {
	v := reflect.ValueOf(input)

	switch v.Kind() {
	case reflect.Float64:
		return true
	case reflect.Float32:
		return true
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
