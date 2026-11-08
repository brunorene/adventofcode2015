package main

func part2(lines []commandLine) int64 {
	lights := [1000][1000]int{}

	for _, cl := range lines {
		switch cl.command {
		case on:
			for x := cl.start.x; x <= cl.end.x; x++ {
				for y := cl.start.y; y <= cl.end.y; y++ {
					lights[x][y] += 1
				}
			}
		case off:
			for x := cl.start.x; x <= cl.end.x; x++ {
				for y := cl.start.y; y <= cl.end.y; y++ {
					lights[x][y] -= 1
					if lights[x][y] < 0 {
						lights[x][y] = 0
					}
				}
			}
		case toggle:
			for x := cl.start.x; x <= cl.end.x; x++ {
				for y := cl.start.y; y <= cl.end.y; y++ {
					lights[x][y] += 2
				}
			}
		}
	}

	var brightness int64

	for _, line := range lights {
		for _, br := range line {
			brightness += int64(br)
		}
	}

	return brightness
}
