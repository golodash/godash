package array

import "errors"

// Gets a slice and size as input and splits the slice into pieces
func Chunk[value any](slice []value, size int) ([][]value, error) {
	var len int = len(slice)
	var lenPieces int

	if size <= 0 {
		return nil, errors.New("size parameter must be more that 0")
	}

	if float32(len) / float32(size) != float32(len / size) {
		lenPieces = (len / size) + 1
	} else {
		lenPieces = len / size
	}

	var a [][]value = make([][]value, lenPieces)
	var i int = size
	var j int = 0

	for ; i < len; i = i + size {
		a[j] = slice[i - size:i]
		j = j + 1
	}
	if len > i - size {
		a[j] = slice[i - size:len]
	}

	return a, nil
}
