package slice

import (
	"errors"

	"github.com/gotorn/godash/internal"
)

// Gets a slice and size as input and splits the slice into pieces in length of the size
func Chunk(slice interface{}, size int) ([][]interface{}, error) {

	s, err := internal.InterfaceToSlice(slice)
	if err != nil {
		return nil, err
	}

	if size <= 0 {
		return nil, errors.New("size must be greater than zero")
	}

	var lenPieces int
	var length int = len(s)

	if float32(length)/float32(size) != float32(length/size) {
		lenPieces = (length / size) + 1
	} else {
		lenPieces = length / size
	}

	var chunks [][]interface{} = make([][]interface{}, lenPieces)
	var i int = size
	var j int = 0

	for ; i < length; i = i + size {
		chunks[j] = s[i-size : i]
		j = j + 1
	}
	if length > i-size {
		chunks[j] = s[i-size : length]
	}

	return chunks, nil

}
