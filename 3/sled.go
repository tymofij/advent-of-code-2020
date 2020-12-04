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
*/
const height = 323

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
	j := 0
	treeCount := 0

	for i := 0; i < height; i++ {
		if slope[i][j] == '#' {
			treeCount++
		}
		j = (j + 3) % width
	}
	fmt.Println(treeCount)
}
