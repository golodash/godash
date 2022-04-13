package slices

import "github.com/golodash/godash/internal"

func Initial(slice interface{}) ([]interface{}, error) {
	s, err := internal.InterfaceToSlice(slice)
	if err != nil {
		return nil, err
	}

	if len(s) > 0 {
		return s[:len(s)-1], nil
	} else {
		return s, nil
	}
}
