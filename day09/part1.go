package main

import (
	"maps"
	"math"
	"slices"
)

func part1(input string) int {
	graph := parse(input)

	cities := slices.Collect(maps.Keys(graph))

	minDist := math.MaxInt

	for _, path := range candidates(cities, len(cities)) {
		curr := distance(path, graph)

		if curr < minDist {
			minDist = curr
		}
	}

	return minDist
}
