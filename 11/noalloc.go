package main

func nextStateNoalloc(seats [][]byte, new [][]byte, tolerance, visibilityLimit int) bool {
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

func occupiedSeatsWhenStabilizedNoalloc(seats [][]byte, tolerance, visibilityLimit int) int {
	new := make([][]byte, len(seats)) // allocate next step just once
	for i := range new {
		new[i] = make([]byte, len(seats[0]))
	}
	prev := make([][]byte, len(seats)) // we will be writing here, do not want to modify the original
	for i := range prev {
		prev[i] = make([]byte, len(seats[0]))
		copy(prev[i], seats[i])
	}

	modified := nextStateNoalloc(prev, new, tolerance, visibilityLimit)
	for modified {
		prev, new = new, prev // updating prev state of seats to newer one, setting up new for reuse
		// show(seats)
		modified = nextStateNoalloc(prev, new, tolerance, visibilityLimit)
	}

	res := 0
	for _, line := range new {
		for _, seat := range line {
			if seat == TAKEN {
				res++
			}
		}
	}
	return res
}
