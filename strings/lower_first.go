package strings

import (
	"strings"
)

// Converts the first character of string to lower case.
//
// Complexity: O(n)
func LowerFirst(input string) string {
	trimInput := strings.Trim(input, " \t\n")
	if len(trimInput) == 0 {
		return trimInput
	}

	outputBuilder := strings.Builder{}
	firstLetter := trimInput[0]
	isCap := firstLetter >= 'A' && firstLetter <= 'Z'

	if isCap {
		firstLetter += 'a'
		firstLetter -= 'A'
	}

	outputBuilder.WriteByte(firstLetter)
	if len(trimInput) > 1 {
		outputBuilder.WriteString(trimInput[1:])
	}

	return outputBuilder.String()
}
