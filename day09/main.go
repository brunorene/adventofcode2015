package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

//go:embed input
var input string

func main() {
	println("part 1:", part1(input))
	println("part 2:", part2(input))
}

func distance(cities []string, graph map[string]map[string]int) (result int) {
	for idx, city := range cities[:len(cities)-1] {
		result += graph[city][cities[idx+1]]
	}

	return result
}

func candidates(cities []string, length int) (result [][]string) {
	for _, city := range cities {
		if length == 1 {
			result = append(result, []string{city})

			continue
		}

		path := candidates(cities, length-1)

		for _, others := range path {
			if slices.Index(others, city) < 0 {
				result = append(result, append([]string{city}, others...))
			}
		}
	}

	return result
}

func parse(input string) map[string]map[string]int {
	scanner := bufio.NewScanner(strings.NewReader(input))

	var city1, city2 string
	var dist int

	graph := make(map[string]map[string]int)

	for scanner.Scan() {
		line := scanner.Text()

		if _, err := fmt.Sscanf(line, "%s to %s = %d", &city1, &city2, &dist); err == nil {
			if _, exists := graph[city1]; !exists {
				graph[city1] = map[string]int{city2: dist}
			}

			if _, exists := graph[city2]; !exists {
				graph[city2] = map[string]int{city1: dist}
			}

			graph[city1][city2] = dist
			graph[city2][city1] = dist

			continue
		}

		panic("invalid line " + line)
	}

	return graph
}
