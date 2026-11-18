package main

import (
	"maps"
	"slices"
)

func part2(input string) int {
	graph := parse(input)

	cities := slices.Collect(maps.Keys(graph))

	var maxDist int

	for _, path := range candidates(cities, len(cities)) {
		curr := distance(path, graph)

		if curr > maxDist {
			maxDist = curr
		}
	}

	return maxDist
}
