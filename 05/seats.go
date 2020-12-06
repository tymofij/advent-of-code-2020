package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	seatIds := make([]int, 0, 1000)
	for scanner.Scan() {
		line := scanner.Text()
		seatID := 0
		for i := 0; i <= 9; i++ {
			seatID <<= 1
			var bit int
			switch line[i] {
			case 'F', 'L':
				bit = 0
			case 'B', 'R':
				bit = 1
			}
			seatID += bit
		}
		seatIds = append(seatIds, seatID)
	}
	sort.Ints(seatIds)
	fmt.Println("Max:", seatIds[len(seatIds)-1])

	for i, v := range seatIds {
		if i > 0 && seatIds[i-1]+1 != v {
			fmt.Println("Missing:", v-1)
		}
	}
}
