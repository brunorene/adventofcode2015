package main

func part1(lines []commandLine) int {
	lights := make(map[point]struct{})

	for _, cl := range lines {
		switch cl.command {
		case on:
			for x := cl.start.x; x <= cl.end.x; x++ {
				for y := cl.start.y; y <= cl.end.y; y++ {
					lights[point{x, y}] = struct{}{}
				}
			}
		case off:
			for x := cl.start.x; x <= cl.end.x; x++ {
				for y := cl.start.y; y <= cl.end.y; y++ {
					delete(lights, point{x, y})
				}
			}
		case toggle:
			for x := cl.start.x; x <= cl.end.x; x++ {
				for y := cl.start.y; y <= cl.end.y; y++ {
					if _, exists := lights[point{x, y}]; exists {
						delete(lights, point{x, y})
					} else {
						lights[point{x, y}] = struct{}{}
					}
				}
			}
		}
	}

	return len(lights)
}
