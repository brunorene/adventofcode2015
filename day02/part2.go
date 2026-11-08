package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

func part2(input string) int {
	var ribbon int

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var l, w, h int
		if _, err := fmt.Sscanf(scanner.Text(), "%dx%dx%d", &l, &w, &h); err != nil {
			fmt.Println("bad line:", err)
			continue
		}

		parts := []int{l, w, h}

		sort.Ints(parts)

		ribbon += 2*parts[0] + 2*parts[1] + l*w*h
	}

	return ribbon
}
