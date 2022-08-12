package strings

import "strings"

// Pads string on the left side if it's shorter than length.
// Padding characters are truncated if they exceed length.
//
// Complexity: O(n)
//
// n = input % length
func PadStart(input string, length int, pattern string) string {
	if len(input) >= length || len(pattern) == 0 {
		return input[len(input)-length:]
	}

	remaining := (length - len(input)) % len(pattern)
	number := (length - len(input)) / len(pattern)
	outputBuilder := strings.Builder{}
	patternBytes := []byte(pattern)
	outputBuilder.Write(patternBytes[len(pattern)-remaining:])
	for i := 0; i < number; i++ {
		outputBuilder.Write(patternBytes)
	}

	outputBuilder.WriteString(input)
	return outputBuilder.String()
}
