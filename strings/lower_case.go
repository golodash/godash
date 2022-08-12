package strings

// Converts string, as space separated words, to lower case.
//
// Complexity: O(n)
func LowerCase(input string) string {
	return CustomDelimitedCase(input, ' ', "", false)
}
