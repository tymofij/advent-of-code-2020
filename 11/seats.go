package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const FLOOR = '.'
const EMPTY = 'L'
const TAKEN = '#'

func countVisible(seats [][]byte, i, j, visibilityRange int) int {
	res := 0
	directions := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, d := range directions {
		ni, nj := i+d[0], j+d[1]
		steps := 1
		for 0 <= ni && ni <= len(seats)-1 &&
			0 <= nj && nj <= len(seats[0])-1 {
			if seats[ni][nj] == EMPTY {
				break
			}
			if seats[ni][nj] == TAKEN {
				res++
				break
			}
			if visibilityRange != -1 && steps >= visibilityRange {
				break
			}
			ni += d[0]
			nj += d[1]
			steps++
		}
	}
	return res
}

func show(seats [][]byte) {
	for _, s := range seats {
		fmt.Println(string(s))
	}
	fmt.Println()
}

func nextState(seats [][]byte, new [][]byte, tolerance, visibilityLimit int) bool {
	modified := false
	rows, cols := len(seats), len(seats[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			new[i][j] = seats[i][j]
			if seats[i][j] == FLOOR {
				continue
			}
			neighbours := countVisible(seats, i, j, visibilityLimit)
			if seats[i][j] == EMPTY && neighbours == 0 {
				new[i][j] = TAKEN
				modified = true
			}
			if seats[i][j] == TAKEN && neighbours >= tolerance {
				new[i][j] = EMPTY
				modified = true
			}
		}
	}
	return modified
}

// avoid memory allocations at the cost of rewriting `seats` contents.
// surprisingly does not yield any speedups
func occupiedSeatsWhenStabilized(seats [][]byte, tolerance, visibilityLimit int) int {
	modified := true
	new := make([][]byte, len(seats)) // allocate next step just once
	for i := range new {
		new[i] = make([]byte, len(seats[0]))
	}

	modified = nextState(seats, new, tolerance, visibilityLimit)
	for modified {
		seats, new = new, seats // updating the state of seats to newer one, setting up new for reuse
		// show(seats)
		modified = nextState(seats, new, tolerance, visibilityLimit)
	}

	res := 0
	for _, line := range seats {
		for _, seat := range line {
			if seat == TAKEN {
				res++
			}
		}
	}
	return res
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	text := strings.TrimSpace(string(data))
	seats := [][]byte{}
	for _, line := range strings.Split(text, "\n") {
		seats = append(seats, []byte(line))
	}
	fmt.Println("part 1:", occupiedSeatsWhenStabilized(seats, 4, 1))

	seats = [][]byte{}
	for _, line := range strings.Split(text, "\n") {
		seats = append(seats, []byte(line))
	}
	fmt.Println("part 2:", occupiedSeatsWhenStabilized(seats, 5, -1))
}
