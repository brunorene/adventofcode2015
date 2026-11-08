package main

import (
	"iter"
	"strings"
)

func nice2(word string) bool {
	var found1 bool

	for pair := range Pair(word) {
		first := strings.Index(word, pair)
		last := strings.LastIndex(word, pair)

		if last-first > 1 {
			found1 = true

			break
		}
	}

	for pos, c := range word {
		if pos < len(word)-2 {
			if c == rune(word[pos+2]) && found1 {
				return true
			}
		}
	}

	return false
}

func Pair(word string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for pos, c := range word {
			if pos < len(word)-1 {
				if !yield(string([]byte{byte(c), word[pos+1]})) {
					return
				}
			}
		}
	}
}
