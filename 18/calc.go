package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func do(a, b int, op rune) int {
	switch op {
	case '+':
		return a + b
	case '*':
		return a * b
	}
	return 0
}

func findMatchingBracketPos(runes []rune, startPos int) int {
	brackets := 1
	for i := startPos + 1; i < len(runes); i++ {
		if runes[i] == ')' {
			brackets--
		}
		if runes[i] == '(' {
			brackets++
		}
		if brackets == 0 {
			return i
		}
	}
	return -1
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	text := strings.TrimSpace(string(data))
	lines := strings.Split(text, "\n")

	s := 0
	for _, line := range lines {
		s += simpleCalc(line)
	}
	fmt.Println(s)
}
