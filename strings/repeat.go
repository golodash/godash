package strings

import "strings"

// Repeats the given string count times.
//
// Complexity: O(n)
//
// n = count
func Repeat(input string, count int) string {
	if input == "" {
		return ""
	}

	var outputBuilder strings.Builder
	for i := 0; i < count; i++ {
		outputBuilder.WriteString(input)
	}

	return outputBuilder.String()
}
