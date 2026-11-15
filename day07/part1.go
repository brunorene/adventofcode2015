package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func part1(wireName string, input string) uint16 {
	cache := make(map[string]uint16)

	return parse(input)[wireName](cache)
}

func parseWire(val string, store map[string]wire) wire {
	num, err := strconv.Atoi(val)
	if err == nil {
		return value(uint16(num))
	}

	return get(val, store)
}

func parse(input string) map[string]wire {
	scanner := bufio.NewScanner(strings.NewReader(input))

	wires := make(map[string]wire)

	var left, right, out string
	var num int

	for scanner.Scan() {
		line := scanner.Text()

		if _, err := fmt.Sscanf(line, "%s RSHIFT %d -> %s", &left, &num, &out); err == nil {
			wires[out] = rShift(parseWire(left, wires), value(uint16(num)))

			continue
		}

		if _, err := fmt.Sscanf(line, "%s LSHIFT %d -> %s", &left, &num, &out); err == nil {
			wires[out] = lShift(parseWire(left, wires), value(uint16(num)))

			continue
		}

		if _, err := fmt.Sscanf(line, "%s AND %s -> %s", &left, &right, &out); err == nil {
			wires[out] = and(parseWire(left, wires), parseWire(right, wires))

			continue
		}

		if _, err := fmt.Sscanf(line, "%s OR %s -> %s", &left, &right, &out); err == nil {
			wires[out] = or(parseWire(left, wires), parseWire(right, wires))

			continue
		}

		if _, err := fmt.Sscanf(line, "NOT %s -> %s", &left, &out); err == nil {
			wires[out] = not(parseWire(left, wires))

			continue
		}

		if _, err := fmt.Sscanf(line, "%d -> %s", &num, &out); err == nil {
			wires[out] = value(uint16(num))

			continue
		}

		if _, err := fmt.Sscanf(line, "%s -> %s", &left, &out); err == nil {
			wires[out] = get(left, wires)

			continue
		}

		panic(line + " was not recognised")
	}

	return wires
}

type wire func(cache map[string]uint16) uint16

func or(left, right wire) wire {
	return func(cache map[string]uint16) uint16 {
		println("or")
		return left(cache) | right(cache)
	}
}

func and(left, right wire) wire {
	return func(cache map[string]uint16) uint16 {
		println("and")
		return left(cache) & right(cache)
	}
}

func rShift(left, shift wire) wire {
	return func(cache map[string]uint16) uint16 {
		println("rshift")
		return left(cache) >> shift(cache)
	}
}

func lShift(left, shift wire) wire {
	return func(cache map[string]uint16) uint16 {
		println("lshift")
		return left(cache) << shift(cache)
	}
}

func not(field wire) wire {
	return func(cache map[string]uint16) uint16 {
		println("not")
		return ^field(cache)
	}
}

func value(num uint16) wire {
	return func(cache map[string]uint16) uint16 {
		println("num", num)
		return num
	}
}

func get(name string, store map[string]wire) wire {
	return func(cache map[string]uint16) uint16 {
		println("get", name)

		if value, exists := cache[name]; exists {
			return value
		}

		value := store[name](cache)

		cache[name] = value

		return value
	}
}
