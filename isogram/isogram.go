// Package isogram provides utility methods for isograms
package isogram

import (
	"unicode"
)

// IsIsogram checks if a string is an isogram
func IsIsogram(input string) bool {
	letters := make(map[rune]bool)
	for _, r := range []rune(input) {
		if unicode.IsLetter(r) {
			r = unicode.ToLower(r)
			_, ok := letters[r]
			if !ok {
				letters[r] = true
			} else {
				return false
			}
		}
	}
	return true
}
