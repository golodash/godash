package strings

import "strings"

// Pads string on the left side if it's shorter than length.
// Padding characters are truncated if they exceed length.
//
// Complexity: O(n)
//
// n = (length - len(input)) / len(pattern)
func PadEnd(input string, length int, pattern string) string {
	if len(input) >= length || len(pattern) == 0 {
		return input[:length]
	}

	remaining := (length - len(input)) % len(pattern)
	number := (length - len(input)) / len(pattern)
	outputBuilder := strings.Builder{}
	patternBytes := []byte(pattern)
	outputBuilder.WriteString(input)
	for i := 0; i < number; i++ {
		outputBuilder.Write(patternBytes)
	}
	if remaining != 0 {
		outputBuilder.Write(patternBytes[:remaining])
	}

	return outputBuilder.String()
}
