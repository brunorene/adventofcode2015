package main

func part2(wireName string, input string) uint16 {
	cache := make(map[string]uint16)

	result := parse(input)[wireName](cache)

	cache = make(map[string]uint16)
	cache["b"] = result

	result = parse(input)[wireName](cache)

	return result
}
