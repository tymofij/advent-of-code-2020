package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type coord struct {
	x, y, z int
}

func neighbours(pos coord, data map[coord]bool) int {
	n := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if dx == 0 && dy == 0 && dz == 0 {
					continue
				}
				has, _ := data[coord{pos.x + dx, pos.y + dy, pos.z + dz}]
				if has {
					n++
				}
			}
		}
	}
	return n
}

var dots = map[coord]bool{}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	text := strings.TrimSpace(string(data))
	for i, line := range strings.Split(text, "\n") {
		for j, char := range []byte(line) {
			if char == '#' {
				dots[coord{i, j, 0}] = true
			}
		}
	}

	for day := 0; day <= 5; day++ {
		minX, minY, minZ := math.MaxInt32, math.MaxInt32, math.MaxInt32
		maxX, maxY, maxZ := math.MinInt32, math.MinInt32, math.MinInt32
		newDots := map[coord]bool{}
		for k := range dots {
			if k.x > maxX {
				maxX = k.x
			}
			if k.y > maxY {
				maxY = k.y
			}
			if k.z > maxZ {
				maxZ = k.z
			}
			if k.x < minX {
				minX = k.x
			}
			if k.y < minY {
				minY = k.y
			}
			if k.z < minZ {
				minZ = k.z
			}
		}
		for x := minX - 1; x <= maxX+1; x++ {
			for y := minY - 1; y <= maxY+1; y++ {
				for z := minZ - 1; z <= maxZ+1; z++ {
					pos := coord{x, y, z}
					active, _ := dots[pos]
					n := neighbours(pos, dots)
					if active == true {
						if n == 2 || n == 3 {
							newDots[pos] = true
						}
					} else {
						if n == 3 {
							newDots[pos] = true
						}
					}
				}
			}
		}
		dots = newDots

		fmt.Println("Day", day, ", active:", len(dots))
	}
}
