package main

func part2(input string) int {
	currentRobot := [2]int{0, 0}
	currentSanta := [2]int{0, 0}
	houses := map[[2]int]struct{}{currentRobot: struct{}{}}

	for pos, move := range input {
		func() {
			current := currentSanta
			if pos%2 == 0 {
				defer func() { currentSanta = current }()
			} else {
				current = currentRobot

				defer func() { currentRobot = current }()
			}

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
		}()
	}

	return len(houses)
}
