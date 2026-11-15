package main

import (
	"bufio"
	"strings"
)

func part2(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var sum int

	for scanner.Scan() {
		line := scanner.Text()

		newCode := strings.ReplaceAll(strings.ReplaceAll(line, `\`, `\\`), `"`, `\"`)

		sum += len(newCode) - len(line) + 2
	}

	return sum
}
