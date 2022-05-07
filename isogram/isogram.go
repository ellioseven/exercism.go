package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	letters := make(map[rune]int)

	for _, letter := range strings.ToLower(word) {
		_letter := unicode.ToLower(letter)

		if !unicode.IsLetter(_letter) {
			continue
		}

		if letters[_letter] == 1 {
			return false
		}

		letters[_letter]++
	}

	return true
}
