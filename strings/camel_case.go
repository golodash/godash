package strings

// Converts string to camel case.
//
// Complexity: O(n)
func CamelCase(input string) string {
	return internalCamelCase(input, false)
}
