package strings

// Converts string to snake case.
//
// Complexity: O(n)
func SnakeCase(input string) string {
	return CustomDelimitedCase(input, '_', "", false)
}
