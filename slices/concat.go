package slices

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/golodash/godash/internal"
)

var errTypeCheck error = errors.New("element in %d index does not match its type with others")

// Creates a new slice concatenating slice with other one.
func Concat(slice interface{}, values ...interface{}) ([]interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	s := reflect.TypeOf(slice)
	v := reflect.ValueOf(slice)

	for i := range values {
		if reflect.TypeOf(values[i]) != nil {
			if reflect.TypeOf(values[i]).Kind() == reflect.Slice && (reflect.TypeOf(values[i]).String() == s.String() || s.Elem().Kind() == reflect.Interface) {
				a := values[i]
				if s.Elem().Kind() == reflect.Interface && reflect.TypeOf(values[i]).Kind() != reflect.Interface {
					a, _ = internal.InterfaceToSlice(values[i])
				}
				v = reflect.AppendSlice(v, reflect.ValueOf(a))
			} else if reflect.TypeOf(values[i]).Kind() == s.Elem().Kind() || s.Elem().Kind() == reflect.Interface {
				v = reflect.Append(v, reflect.ValueOf(values[i]))
			} else {
				return nil, fmt.Errorf(errTypeCheck.Error(), i)
			}
		}
	}

	res, err := internal.InterfaceToSlice(v.Interface())
	if err != nil {
		return nil, err
	}

	return res, nil
}
