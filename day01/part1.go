package main

func part1(input string) int {
	var floor int

	for _, c := range input {
		if c == '(' {
			floor++
		}

		if c == ')' {
			floor--
		}
	}

	return floor
}
