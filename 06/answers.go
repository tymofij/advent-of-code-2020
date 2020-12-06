package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	text, _ := ioutil.ReadFile("input.txt")
	groups := strings.Split(string(text), "\n\n")
	sumUniques, sumIntersecting := 0, 0

	for _, group := range groups {
		counts := make(map[rune]int, 30)
		answers := strings.Fields(group)
		groupSize := len(answers)

		for _, answer := range answers {
			for _, char := range answer {
				counts[char]++
			}
		}
		sumUniques += len(counts)

		intersecting := 0
		for char := range counts {
			if counts[char] == groupSize {
				intersecting++
			}
		}
		sumIntersecting += intersecting
	}
	fmt.Println(sumUniques)
	fmt.Println(sumIntersecting)
}
