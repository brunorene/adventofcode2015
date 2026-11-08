package main

func part1(input string) int {
	current := [2]int{0, 0}
	houses := map[[2]int]struct{}{current: struct{}{}}

	for _, move := range input {
		switch move {
		case '<':
			current = [2]int{current[0] - 1, current[1]}
		case '>':
			current = [2]int{current[0] + 1, current[1]}
		case '^':
			current = [2]int{current[0], current[1] - 1}
		case 'v':
			current = [2]int{current[0], current[1] + 1}
		}

		houses[current] = struct{}{}
	}

	return len(houses)
}
