package slices

import (
	"fmt"
	"reflect"

	"github.com/golodash/godash/internal"
)

func Difference(slice interface{}, notIncluded interface{}) ([]interface{}, error) {
	if err := internal.AreComparable(slice, notIncluded); err != nil {
		return nil, err
	}
	if err1, err2 := internal.SliceCheck(slice), internal.SliceCheck(notIncluded); err1 != nil || err2 != nil {
		if err2 != nil {
			return nil, err2
		}
		return nil, err1
	}

	s := reflect.ValueOf(slice)
	notIn := reflect.ValueOf(notIncluded)

	for i := s.Len() - 1; i > -1; i-- {
		if i >= s.Len() {
			continue
		}
	firstLoop:
		for j := 0; j < notIn.Len(); j++ {
			res, err := same(s.Index(i), notIn.Index(j))
			if err != nil {
				return nil, err
			}
			if res {
				if i != 0 && i+1 < s.Len() {
					s = reflect.AppendSlice(s.Slice(0, i), s.Slice(i+1, s.Len()))
				} else if i == 0 {
					s = s.Slice(i+1, s.Len())
				} else if i+1 >= s.Len() {
					s = s.Slice(0, i)
				}
				i++
				break firstLoop
			}
		}
	}

	s1, err := internal.InterfaceToSlice(s.Interface())
	if err != nil {
		return nil, err
	}

	return s1, nil
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
