package slices

import (
	"reflect"
	"sort"

	"github.com/golodash/godash/internal"
)

// Removes elements from slice corresponding to indexes and returns a slice of
// remaining elements and removed elements.
//
// Note: Duplicate key numbers in 'remSlice' variable will get removed
//
// Complexity: O(n)
//
// n = length of 'remSlice'
func PullAt(slice interface{}, remSlice []int) (interface{}, interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, nil, err
	}

	sort.Ints(remSlice)
	remSlice = internal.UniqueInt(remSlice)
	sliceValue := reflect.ValueOf(slice)
	removed := reflect.MakeSlice(sliceValue.Type(), 0, sliceValue.Len())
	for i := 0; i < len(remSlice); i++ {
		if !(remSlice[i]-i < sliceValue.Len()) {
			break
		}
		removed = reflect.Append(removed, reflect.ValueOf(sliceValue.Index(remSlice[i]-i).Interface()))
		sliceValue = reflect.AppendSlice(sliceValue.Slice(0, remSlice[i]-i), sliceValue.Slice(remSlice[i]+1-i, sliceValue.Len()))
	}

	return sliceValue.Interface(), removed.Interface(), nil
}
