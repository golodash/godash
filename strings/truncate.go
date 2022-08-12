package strings

import (
	"strings"

	"github.com/golodash/godash/internal"
)

// Truncates string if it's longer than the given maximum string length.
// The last characters of the truncated string are replaced with the
// omission string.
//
// Complexity: O(1)
func Truncate(input string, length int, separators []rune, omission string) string {
	tempInput := strings.TrimSpace(input)
	if len(tempInput) <= length {
		return tempInput
	}

	tempInput = tempInput[:length]
	i := length - 1
	for ; i > -1; i-- {
		if internal.CustomIsSeparator(rune(tempInput[i]), separators) {
			tempInput = tempInput[:i]
			break
		}
	}

	return strings.TrimSpace(tempInput) + omission
}
