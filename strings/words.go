package strings

import (
	"strings"

	"github.com/golodash/godash/internal"
)

var wordsSeparators = []rune{',', '.', '/', '\\', '\'', '"', ':', ';', ']', '[', '=', '+', '-', ')', '(', '*', '&', '^', '%', '$', '#', '@', '!', '~', '`', '|', '?', ' ', '\t', '\n', '_', '<', '>', '}', '{'}

// Splits string into a slice of its words.
//
// Complexity: O(n)
func Words(input string) []string {
	outputSlice := []string{}
	sawOneSeparatorLastTime := false
	wordBuilder := strings.Builder{}
	for _, letter := range []byte(input) {
		if internal.CustomIsSeparator(rune(letter), wordsSeparators) {
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
