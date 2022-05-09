package slices

import (
	"github.com/golodash/godash/internal"
)

//This function creates a duplicate-free version of an slice.
//This method is designed and optimized for sorted slices.
func SortedUnique(slice interface{}) (interface{}, error) {
	err := internal.SliceCheck(slice)
	if err != nil {
		return nil, err
	}
	n, err := internal.InterfaceToSlice(slice)
	if err != nil {
		return nil, err
	}

	m := make(map[interface{}]bool)
	var unique []interface{}

	for _, value := range n {
		if _, ok := m[value]; !ok {
			m[value] = true
			unique = append(unique, value)
		}
	}
	return unique, nil
}
