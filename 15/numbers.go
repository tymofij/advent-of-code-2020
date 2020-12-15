package main

import (
	"fmt"
)

func main() {
	data := []int{10, 16, 6, 0, 1, 17}
	lastSpoken := map[int]int{}
	for k, v := range data {
		if k < len(data)-1 {
			lastSpoken[v] = k + 1
		}
	}
	prev := data[len(data)-1]
	turn := len(data)

	var next int
	for turn < 30_000_000 {
		prevTurn, ok := lastSpoken[prev]
		if ok {
			next = turn - prevTurn
		} else {
			next = 0
		}
		lastSpoken[prev] = turn
		prev = next
		turn++
	}

	fmt.Println(turn, prev)
}
