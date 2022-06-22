package slices

// Gets the first element of slice.
//
// Complexity: O(1)
func First(slice interface{}) interface{} {
	return Head(slice)
}
