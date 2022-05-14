package slices

import (
	"fmt"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Returns a slice of slice values not included in the
// other given slice using equality comparisons.
//
// Note: In comparing fields of a struct, unexported fields
// are ignored.
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
			res, err := same(s[i], notIn.Index(j).Interface())
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

func same(value1 interface{}, value2 interface{}) (condition bool, err error) {
	condition, err = true, nil
	v1 := reflect.ValueOf(value1)
	v2 := reflect.ValueOf(value2)

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
		for i, j := 0, 0; i < v1.Len(); i, j = i+1, j+1 {
			condition, err = same(v1.Index(i).Interface(), v2.Index(j).Interface())
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

		keys1 := v1.MapKeys()
		keys2 := v2.MapKeys()
		if len(keys1) != len(keys2) {
			condition, err = false, nil
			return
		}

		for i := 0; i < len(keys1); i = i + 1 {
			condition, err = same(v1.MapIndex(keys1[i]).Interface(), v2.MapIndex(keys1[i]).Interface())
			if err != nil || !condition || !keyIsInHereToo(keys1[i], keys2) {
				condition, err = false, nil
				return
			}
		}
	case reflect.Struct:
		condition, err = value1 == value2, nil
	case reflect.Ptr:
		condition, err = same(v1.Elem().Interface(), v2.Elem().Interface())
	default:
		if v1.Interface() != v2.Interface() {
			condition, err = false, nil
		}
	}

	return
}

func Without(slice interface{}, notIncluded interface{}) ([]interface{}, error) {
	return Difference(slice, notIncluded)
}
