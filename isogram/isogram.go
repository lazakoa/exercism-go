package isogram

import "strings"

func IsIsogram(word string) bool {
	runes := []rune(strings.ToUpper(word))
	counts := make(map[rune]int)

	for _, r := range runes {
		if InRange(r) {
			counts[r]++
		}
	}

	for _, count := range counts {
		if count > 1 {
			return false
		}
	}
	return true
}

func InRange(letter rune) bool {
	return (97 <= letter && letter <= 122) || (65 <= letter && letter <= 90)
}
