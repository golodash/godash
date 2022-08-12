package strings

import (
	"strings"
)

// Converts the first character of string to upper case.
//
// Complexity: O(n)
func UpperFirst(input string) string {
	trimInput := strings.TrimSpace(input)
	if len(trimInput) == 0 {
		return trimInput
	}

	outputBuilder := strings.Builder{}
	firstLetter := trimInput[0]
	isLow := firstLetter >= 'a' && firstLetter <= 'z'

	if isLow {
		firstLetter += 'A'
		firstLetter -= 'a'
	}

	outputBuilder.WriteByte(firstLetter)
	if len(trimInput) > 1 {
		outputBuilder.WriteString(trimInput[1:])
	}

	return outputBuilder.String()
}
