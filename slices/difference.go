package slices

import (
	"fmt"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Returns an slice of slice values not included in the
// other given slice using SameValueZero `and` SameValueNonNumber
// for equality comparisons.
//
// Link to descriptions of SameValueZero = https://262.ecma-international.org/7.0/#sec-samevaluezero
//
// Link to descriptions of SameValueNonNumber = https://262.ecma-international.org/7.0/#sec-samevaluenonnumber
func Difference(slice interface{}, notIncluded interface{}) ([]interface{}, error) {
	if err := internal.CheckSameType(slice, notIncluded); err != nil {
		return nil, err
	}
	if err1, err2 := internal.SliceCheck(slice), internal.SliceCheck(notIncluded); err1 != nil || err2 != nil {
		if err2 != nil {
			return nil, err2
		}
		return nil, err1
	}

	s, err := internal.InterfaceToSlice(slice)
	if err != nil {
		return nil, err
	}

	notIn := reflect.ValueOf(notIncluded)

	for i := len(s) - 1; i > -1; i-- {
		if i >= len(s) {
			continue
		}
	firstLoop:
		for j := 0; j < notIn.Len(); j++ {
			res, err := same(reflect.ValueOf(s[i]), notIn.Index(j))
			if err != nil {
				return nil, err
			}
			if res {
				if i != 0 && i+1 < len(s) {
					s = append(s[0:i], s[i+1:]...)
				} else if i == 0 {
					s = s[i+1:]
				} else if i+1 >= len(s) {
					s = s[0:i]
				}
				i++
				break firstLoop
			}
		}
	}

	return s, nil
}

func keyIsInHereToo(key reflect.Value, keys []reflect.Value) bool {
	for i := range keys {
		if key.Interface() == keys[i].Interface() {
			return true
		}
	}

	return false
}

func same(v1 reflect.Value, v2 reflect.Value) (condition bool, err error) {
	condition, err = true, nil

	if v1.Kind() != v2.Kind() {
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
		for i, j := 0, 0; i < v1.Len(); i, j = i+1, j+1 {
			condition, err = same(v1.Index(i), v2.Index(j))
			if err != nil || !condition {
				return
			}
		}
	case reflect.Map:
		if v1.Len() != v2.Len() {
			condition, err = false, nil
			return
		}

		keys1 := v1.MapKeys()
		keys2 := v2.MapKeys()
		if len(keys1) != len(keys2) {
			condition, err = false, nil
			return
		}

		for i := 0; i < len(keys1); i = i + 1 {
			condition, err = same(v1.MapIndex(keys1[i]), v2.MapIndex(keys1[i]))
			if err != nil || !condition || !keyIsInHereToo(keys1[i], keys2) {
				return
			}
		}
	case reflect.Struct:
		if reflect.TypeOf(v1) != reflect.TypeOf(v2) {
			condition, err = false, nil
			return
		}

		for i := 0; i < v1.NumField(); i++ {
			condition, err = same(v1.Field(i), v2.Field(i))
			if err != nil || !condition {
				return
			}
		}
	case reflect.Ptr:
		condition, err = same(v1.Elem(), v2.Elem())
	default:
		if v1.Interface() != v2.Interface() {
			condition, err = false, nil
		}
	}

	return
}
