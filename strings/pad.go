package strings

import (
	"strings"
)

// Pads string on the left and right sides if it's shorter than length.
// Padding characters are truncated if they can't be evenly divided by length.
//
// Complexity: O(n)
//
// n = (length - len(input)) / len(pattern)
func Pad(input string, length int, pattern string) string {
	if len(input) >= length || len(pattern) == 0 {
		halfRemainingLength := (len(input) - length) / 2
		return input[halfRemainingLength+((len(input)-length)%2) : len(input)-halfRemainingLength]
	}

	lengthInputDifference := (length - len(input))
	number := (lengthInputDifference / len(pattern)) / 2
	outputBuilder := strings.Builder{}
	patternBuilder := strings.Builder{}
	patternBytes := []byte(pattern)
	for i := 0; i < number; i++ {
		patternBuilder.Write(patternBytes)
	}

	halfRemaining := (lengthInputDifference - patternBuilder.Len()*2) / 2
	RestOfHalfRemaining := (lengthInputDifference - patternBuilder.Len()*2) % 2
	outputBuilder.WriteString(pattern[len(pattern)-halfRemaining-RestOfHalfRemaining:] + patternBuilder.String())
	outputBuilder.WriteString(input)
	outputBuilder.WriteString(patternBuilder.String() + pattern[:halfRemaining])

	return outputBuilder.String()
}
