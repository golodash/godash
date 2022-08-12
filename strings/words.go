package strings

import "strings"

// Splits string into an array of its words.
//
// Complexity: O(n)
func Words(input string) []string {
	outputSlice := []string{}
	sawOneSeparatorLastTime := false
	wordBuilder := strings.Builder{}
	for _, letter := range []byte(input) {
		if isSeparator(letter) {
			if !sawOneSeparatorLastTime && wordBuilder.Len() != 0 {
				outputSlice = append(outputSlice, wordBuilder.String())
				wordBuilder.Reset()
			}
			sawOneSeparatorLastTime = true
		} else {
			wordBuilder.WriteByte(letter)
			sawOneSeparatorLastTime = false
		}
	}

	if wordBuilder.Len() != 0 {
		outputSlice = append(outputSlice, wordBuilder.String())
	}

	return outputSlice
}

var separators = []uint8{',', '.', '/', '\\', '\'', '"', ':', ';', ']', '[', '=', '+', '-', ')', '(', '*', '&', '^', '%', '$', '#', '@', '!', '~', '`', '|', '?', ' ', '\t', '\n', '_', '<', '>', '}', '{'}

func isSeparator(letter uint8) bool {
	for _, separator := range separators {
		if separator == letter {
			return true
		}
	}
	return false
}
