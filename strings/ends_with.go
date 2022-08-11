package strings

// Checks if string ends with the given target string.
//
// Complexity: O(1)
func EndsWith(input string, end string) bool {
	if len(input) < len(end) {
		return false
	}

	return input[len(input)-len(end):] == end
}
