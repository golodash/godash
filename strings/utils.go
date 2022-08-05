/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2015 Ian Coleman
 * Copyright (c) 2018 Ma_124, <github.com/Ma124>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, Subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or Substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package strings

import "strings"

// CustomDelimitedCase converts a string to SCREAMING.DELIMITED.SNAKE.CASE if 'delimiter'
// equals to '.' and 'screaming' equals to true or to delimited.snake.case if 'delimiter'
// equals to '.' and screaming equals to false.
//
// 'ignore' variable is what you want the function to ignore and do not change.
//
// Credits: https://github.com/iancoleman/strcase/blob/master/snake.go#L66
//
// Complexity: O(n)
func CustomDelimitedCase(input string, delimiter uint8, ignore string, screaming bool) string {
	input = strings.TrimSpace(input)
	n := strings.Builder{}
	n.Grow(len(input) + 2) // nominal 2 bytes of extra space for inserted delimiters
	for i, v := range []byte(input) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if vIsLow && screaming {
			v += 'A'
			v -= 'a'
		} else if vIsCap && !screaming {
			v += 'a'
			v -= 'A'
		}

		// treat acronyms as words, eg for JSONData -> JSON is a whole word
		if i+1 < len(input) {
			next := input[i+1]
			vIsNum := v >= '0' && v <= '9'
			nextIsCap := next >= 'A' && next <= 'Z'
			nextIsLow := next >= 'a' && next <= 'z'
			nextIsNum := next >= '0' && next <= '9'
			// add underscore if next letter case type is changed
			if (vIsCap && (nextIsLow || nextIsNum)) || (vIsLow && (nextIsCap || nextIsNum)) || (vIsNum && (nextIsCap || nextIsLow)) {
				prevIgnore := ignore != "" && i > 0 && strings.ContainsAny(string(input[i-1]), ignore)
				if !prevIgnore {
					if vIsCap && nextIsLow {
						if prevIsCap := i > 0 && input[i-1] >= 'A' && input[i-1] <= 'Z'; prevIsCap {
							n.WriteByte(delimiter)
						}
					}
					n.WriteByte(v)
					if vIsLow || vIsNum || nextIsNum {
						n.WriteByte(delimiter)
					}
					continue
				}
			}
		}

		if (v == ' ' || v == '_' || v == '-' || v == '.') && !strings.ContainsAny(string(v), ignore) {
			// replace space/underscore/hyphen/dot with delimiter
			n.WriteByte(delimiter)
		} else {
			n.WriteByte(v)
		}
	}

	return n.String()
}

// Converts a string to CamelCase.
//
// Credits: https://github.com/iancoleman/strcase/blob/master/camel.go#L32
//
// Complexity: O(n)
func internalCamelCase(s string, initCase bool) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}

	var isSpecialDelimiter = func(b byte) bool {
		return (b == '_' || b == ' ' || b == '-' || b == '.')
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := initCase
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		}

		if vIsCap || vIsLow {
			n.WriteByte(v)
			capNext = false
		} else if vIsNum := v >= '0' && v <= '9'; vIsNum {
			n.WriteByte(v)
			capNext = true
		} else if !isSpecialDelimiter(v) {
			n.WriteByte(v)
			capNext = false
		} else {
			capNext = isSpecialDelimiter(v)
		}
	}

	return n.String()
}
