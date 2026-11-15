package main

import (
	"bufio"
	"regexp"
	"strings"
)

func part1(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var sum int

	for scanner.Scan() {
		line := scanner.Text()

		sum += len(line) - codeLen(line)
	}

	return sum
}

func codeLen(code string) (count int) {
	code = code[1 : len(code)-1]

	hexa := regexp.MustCompile(`\\x[0-9a-f]{2}`)

	for {
		lastLen := len(code)

		code = strings.Replace(code, `\\`, "", 1)

		if len(code)+2 == lastLen {
			count++

			continue
		}

		code = strings.Replace(code, `\"`, "", 1)

		if len(code)+2 == lastLen {
			count++

			continue
		}

		match := hexa.FindString(code)

		if match != "" {
			count++

			code = strings.Replace(code, match, "", 1)

			continue
		}

		break
	}

	return len(code) + count
}
