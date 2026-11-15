package main

import _ "embed"

//go:embed input
var input string

func main() {
	println("part 1:", part1(input))
	println("part 2:", part2(input))
}
