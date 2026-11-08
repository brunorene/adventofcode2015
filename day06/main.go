package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input
var input string

func main() {
	println("part 1:", part1(parse(input)))
	println("part 2:", part2(parse(input)))
}

type command int

const (
	on command = iota
	off
	toggle
)

type point struct {
	x, y int
}

type commandLine struct {
	command command
	start   point
	end     point
}

func parse(input string) (lines []commandLine) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	lines = make([]commandLine, 0, strings.Count(input, "\n"))

	for scanner.Scan() {
		line := scanner.Text()

		var start, end point

		_, errToggle := fmt.Sscanf(line, "toggle %d,%d through %d,%d", &start.x, &start.y, &end.x, &end.y)
		if errToggle == nil {
			lines = append(lines, commandLine{
				command: toggle,
				start:   start,
				end:     end,
			})

			continue
		}

		_, errOn := fmt.Sscanf(line, "turn on %d,%d through %d,%d", &start.x, &start.y, &end.x, &end.y)
		if errOn == nil {
			lines = append(lines, commandLine{
				command: on,
				start:   start,
				end:     end,
			})

			continue
		}

		_, errOff := fmt.Sscanf(line, "turn off %d,%d through %d,%d", &start.x, &start.y, &end.x, &end.y)
		if errOff == nil {
			lines = append(lines, commandLine{
				command: off,
				start:   start,
				end:     end,
			})

			continue
		}
	}

	return lines
}
