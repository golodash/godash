package slices

import (
	"errors"
	"reflect"

	"github.com/golodash/godash/internal"
)

// Gets a slice and a size as input and splits the slice
// into pieces in length of the size.
//
// Complexity: O(n)
func Chunk(slice interface{}, size int) (interface{}, error) {
	if err := internal.SliceCheck(slice); err != nil {
		return nil, err
	}

	if size <= 0 {
		return nil, errors.New("size must be greater than zero")
	}

	var sliceValue = reflect.ValueOf(slice)
	var lenPieces int
	var length int = sliceValue.Len()
	if float32(length)/float32(size) != float32(length/size) {
		lenPieces = (length / size) + 1
	} else {
		lenPieces = length / size
	}

	var typeOfSlice = reflect.SliceOf(reflect.TypeOf(slice))
	var chunks = reflect.MakeSlice(typeOfSlice, 0, lenPieces)
	var i int = size
	var j int = 0
	for ; i < length; i = i + size {
		chunks = reflect.Append(chunks, sliceValue.Slice(i-size, i))
		j = j + 1
	}
	if length > i-size {
		chunks = reflect.Append(chunks, sliceValue.Slice(i-size, length))
	}

	return chunks.Interface(), nil
}
