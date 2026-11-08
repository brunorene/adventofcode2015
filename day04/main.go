package main

import (
	"crypto/md5"
	_ "embed"
	"fmt"
	"strings"
)

func main() {
	println("part 1:", find("yzbqklnj", 5))
	println("part 2:", find("yzbqklnj", 6))
}

func find(input string, zeroes int) int {
	suffix := 1
	for {
		key := fmt.Sprintf("%s%d", input, suffix)

		hash := md5.Sum([]byte(key))
		hexSum := fmt.Sprintf("%x", hash)
		if strings.HasPrefix(hexSum, strings.Repeat("0", zeroes)) {
			return suffix
		}

		suffix++
	}
}
