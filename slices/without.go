package slices

// Returns a slice of 'slice' elements that are not included in the
// other given slice using equality comparisons.
//
// Complexity: O(n*m)
//
// n = length of 'slice'
//
// m = length of 'notIncluded'
func Without(slice, notIncluded interface{}) interface{} {
	return Difference(slice, notIncluded)
}
