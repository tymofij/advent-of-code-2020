package main

import (
	"io/ioutil"
	"strings"
)

// Day5 returns max seat ID for binary space partitioned seats.
func Day5(seats []string) uint {
	max := uint(0)
	for _, seat := range seats {
		n := uint(0)
		for i := range seat {
			n *= 2
			n += ^((uint(seat[i])) >> 2) & 1
		}
		if n > max {
			max = n
		}
	}
	return max
}

func main() {
	text, _ := ioutil.ReadFile("input.txt")
	seats := strings.Fields(string(text))

	print(Day5(seats))
}
