package slices

// This method is like IndexOf except that it iterates over elements of
// 'slice' from right to left. If 'from' is negative, it's used as the offset
// from the end of slice.
//
// Complexity: O(n)
func LastIndexOf(slice, value interface{}, from int) int {
	return indexOf(slice, value, from, false)
}
