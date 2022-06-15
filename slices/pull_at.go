package slices

import (
	"errors"
	"sort"

	"github.com/golodash/godash/internal"
)

// Removes elements from slice corresponding to
// indexes and returns a slice of remaining elements
// and removed elements.
//
// Note: Duplicate keys in remSlice will remove
func PullAt(slice interface{}, remSlice []int) ([]interface{}, []interface{}, error) {
	s, err := internal.InterfaceToSlice(slice)
	if err != nil {
		return nil, nil, errors.New("'slice' variable is not slice type")
	}

	sort.Ints(remSlice)
	remSlice = internal.UniqueInt(remSlice)

	removed := []interface{}{}
	for i := 0; i < len(remSlice); i++ {
		if !(remSlice[i]-i < len(s)) {
			break
		}
		removed = append(removed, s[remSlice[i]-i])
		s = append(s[:remSlice[i]-i], s[remSlice[i]+1-i:]...)
	}

	return s, removed, nil
}
