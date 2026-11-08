package main

import (
	"regexp"
	"strings"
)

func nice1(word string) bool {
	vowels := regexp.MustCompile("[^aeiou]")
	onlyVowels := vowels.ReplaceAllString(word, "")
	if len(onlyVowels) < 3 {
		return false
	}

	noOccur := strings.Index(word, "ab") +
		strings.Index(word, "cd") +
		strings.Index(word, "pq") +
		strings.Index(word, "xy")

	for pos, c := range word {
		if pos < len(word)-1 {
			if c == rune(word[pos+1]) && noOccur == -4 {
				return true
			}
		}
	}

	return false
}
