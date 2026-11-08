package main

import (
	"bufio"
	_ "embed"
	"strings"
)

//go:embed input
var input string

func main() {
	println("part 1:", find(input, nice1))
	println("part 2:", find(input, nice2))
}

func find(input string, niceFunc func(string) bool) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var count int

	for scanner.Scan() {
		line := scanner.Text()

		if niceFunc(line) {
			count++
		}
	}

	return count
}
