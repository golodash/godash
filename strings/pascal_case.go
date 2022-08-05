package strings

// Converts string to pascal case.
//
// Complexity: O(n)
func PascalCase(input string) string {
	return internalCamelCase(input, true)
}
