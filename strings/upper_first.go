package strings

import (
	"strings"
)

// Converts the first character of string to upper case.
//
// Complexity: O(n)
func UpperFirst(input string) string {
	trimesInput := strings.Trim(input, " \t\n")
	if len(trimesInput) == 0 {
		return trimesInput
	}

	outputBuilder := strings.Builder{}
	firstLetter := trimesInput[0]
	isLow := firstLetter >= 'a' && firstLetter <= 'z'

	if isLow {
		firstLetter += 'A'
		firstLetter -= 'a'
	}

	outputBuilder.WriteByte(firstLetter)
	if len(trimesInput) > 1 {
		outputBuilder.WriteString(trimesInput[1:])
	}

	return outputBuilder.String()
}
