package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
You start on the open square (.) in the top-left corner and need to reach the bottom (below the bottom-most row on your map).

These aren't the only trees, the pattern in input.txt repeats to the right many times:

The toboggan can only follow a few specific slopes (you opted for a cheaper model that prefers rational numbers);
start by counting all the trees you would encounter for the slope right 3, down 1:

From your starting position at the top-left, check the position that is right 3 and down 1.
Then, check the position that is right 3 and down 1 from there, and so on until you go past the bottom of the map.

The locations you'd check in the above example are marked here with O where there was an open square and X where there was a tree:

..##.........##.........##.........##.........##.........##.......  --->
#..O#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
.#....X..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#O#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..X...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.X#.......#.##.......#.##.......#.##.......#.##.....  --->
.#.#.#....#.#.#.#.O..#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........X.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.X#...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...#X....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...X.#.#..#...#.#.#..#...#.#.#..#...#.#  --->

Count the number of trees encounted in such diagonal manner.

Part 2:

Determine the number of trees you would encounter if, for each of the following slopes, you start at the top-left corner and traverse the map all the way to the bottom:

    Right 1, down 1.
    Right 3, down 1. (This is the slope you already checked.)
    Right 5, down 1.
    Right 7, down 1.
    Right 1, down 2.


*/
const height = 323 // determined by a fair dice roll

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	slope := make([]string, 0, height)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		slope = append(slope, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	width := len(slope[0])

	totalTreeCount := 1

	steps := [][]int{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}

	for _, stepsPair := range steps {
		stepI, stepJ := stepsPair[0], stepsPair[1]
		curTreeCount := 0
		j := 0
		for i := 0; i < height; i += stepI {
			if slope[i][j] == '#' {
				curTreeCount++
			}
			j = (j + stepJ) % width
		}
		totalTreeCount *= curTreeCount
	}
	fmt.Println(totalTreeCount)
}
