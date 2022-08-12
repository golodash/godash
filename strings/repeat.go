package strings

import "strings"

// Repeats the given string count times.
//
// Complexity: O(n)
//
// n = count
func Repeat(input string, count int) string {
	if len(input) == 0 || count < 1 {
		return ""
	}

	outputBuilder := strings.Builder{}
	outputBuilder.Grow(count * len(input))
	for i := 0; i < count; i++ {
		outputBuilder.WriteString(input)
	}

	return outputBuilder.String()
}
