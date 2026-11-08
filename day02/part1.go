package main

import (
	"bufio"
	"fmt"
	"strings"
)

func part1(input string) int {
	var area int

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var l, w, h int
		if _, err := fmt.Sscanf(scanner.Text(), "%dx%dx%d", &l, &w, &h); err != nil {
			fmt.Println("bad line:", err)
			continue
		}

		area += 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
	}

	return area
}
