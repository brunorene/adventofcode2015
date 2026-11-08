package main

func part2(input string) int {
	var floor int

	for pos, c := range input {
		if c == '(' {
			floor++
		}

		if c == ')' {
			floor--
		}

		if floor < 0 {
			return pos + 1
		}
	}

	return -1
}
