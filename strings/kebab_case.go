package strings

// Converts string to kebab case.
//
// Complexity: O(n)
func KebabCase(input string) string {
	return CustomDelimitedCase(input, '-', "", false)
}
