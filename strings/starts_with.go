package strings

// Checks if string starts with the given target string.
//
// Complexity: O(1)
func StartsWith(input string, start string) bool {
	if len(input) < len(start) {
		return false
	}

	return input[:len(start)] == start
}
