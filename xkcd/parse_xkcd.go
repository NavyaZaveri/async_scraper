package xkcd

import (
	"strings"
	"unicode"
)

func Parse(s string) string {
	wordList := strings.FieldsFunc(s, func(char rune) bool {
		if !unicode.IsLetter(char) {
			return true
		}
		return false
	})
	return strings.Join(wordList, " ")
}
