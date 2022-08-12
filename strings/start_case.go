package strings

import "strings"

// Converts string to Start Case.
//
// Complexity: O(n)
func StartCase(input string) string {
	return startCase(CustomDelimitedCase(input, ' ', "", false))
}

func startCase(input string) string {
	upper := true
	outputBuilder := strings.Builder{}
	outputBuilder.Grow(len(input))
	for _, letter := range []byte(input) {
		isLow := letter >= 'a' && letter <= 'z'
		isSpace := letter == ' '

		if isLow && upper {
			letter += 'A'
			letter -= 'a'
		}
		upper = false

		if isSpace {
			upper = true
		}

		outputBuilder.WriteByte(letter)
	}

	return outputBuilder.String()
}
