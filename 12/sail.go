package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type instruction struct {
	rule byte
	n    int
}

var rules = []instruction{}

func getFinalManhattanDistance(data []instruction) int {
	var x, y float64
	var angle float64 // in degrees
	for _, instr := range data {
		n := float64(instr.n)
		switch instr.rule {
		case 'N':
			y += n
		case 'S':
			y -= n
		case 'E':
			x += n
		case 'W':
			x -= n
		case 'L':
			angle += n
		case 'R':
			angle -= n
		case 'F':
			rad := (math.Pi * angle) / 180
			x += n * math.Cos(rad)
			y += n * math.Sin(rad)
		}
	}
	return int(math.Round(math.Abs(x) + math.Abs(y)))
}

func getFinalManhattanDistanceWithWaypoint(data []instruction) int {
	var x, y float64
	var dx, dy float64 = 10, 1
	for _, instr := range data {
		n := float64(instr.n)
		switch instr.rule {
		case 'N':
			dy += n
		case 'S':
			dy -= n
		case 'E':
			dx += n
		case 'W':
			dx -= n
		case 'F':
			x += n * dx
			y += n * dy
		case 'L', 'R':
			turnAngle := (math.Pi * n) / 180
			radius := math.Sqrt(dx*dx + dy*dy)
			if instr.rule == 'R' {
				turnAngle *= -1
			}
			rad := math.Atan2(dy, dx) + turnAngle
			dy = radius * math.Sin(rad)
			dx = radius * math.Cos(rad)
		}
	}
	return int(math.Round(math.Abs(x) + math.Abs(y)))
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	text := strings.TrimSpace(string(data))
	for _, line := range strings.Split(text, "\n") {
		rule := line[0]
		n, _ := strconv.Atoi(line[1:])
		rules = append(rules, instruction{rule, n})
	}

	fmt.Println("part 1:", getFinalManhattanDistance(rules))
	fmt.Println("part 2:", getFinalManhattanDistanceWithWaypoint(rules))
}
