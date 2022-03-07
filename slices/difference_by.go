package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

func DifferenceBy(slice interface{}, notIncluded interface{}, function interface{}) ([]interface{}, error) {
	if err := internal.AreComparable(slice, notIncluded); err != nil {
		return nil, err
	}
	if err1, err2 := internal.SliceCheck(slice), internal.SliceCheck(notIncluded); err1 != nil || err2 != nil {
		if err2 != nil {
			return nil, err2
		}
		return nil, err1
	}
	if reflect.TypeOf(function).Kind() != reflect.Func {
		return nil, errors.New("`function` variable is not a function")
	}

	functionValue := reflect.ValueOf(function)
	s := reflect.ValueOf(slice)
	notIn := reflect.ValueOf(notIncluded)

	for i := s.Len() - 1; i > -1; i-- {
		if i >= s.Len() {
			continue
		}
	firstLoop:
		for j := 0; j < notIn.Len(); j++ {
			res := functionValue.Call([]reflect.Value{s.Index(i), notIn.Index(j)})

			if res[0].Bool() {
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
