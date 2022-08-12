package strings

import (
	"strings"
)

// Converts the first character of string to lower case.
//
// Complexity: O(n)
func LowerFirst(input string) string {
	trimesInput := strings.Trim(input, " \t\n")
	if len(trimesInput) == 0 {
		return trimesInput
	}

	outputBuilder := strings.Builder{}
	firstLetter := trimesInput[0]
	isCap := firstLetter >= 'A' && firstLetter <= 'Z'

	if isCap {
		firstLetter += 'a'
		firstLetter -= 'A'
	}

	outputBuilder.WriteByte(firstLetter)
	if len(trimesInput) > 1 {
		outputBuilder.WriteString(trimesInput[1:])
	}

	return outputBuilder.String()
}
