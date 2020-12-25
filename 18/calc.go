package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
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

func calc(s string) int {
	runes := []rune(s)
	res := 0
	op := '+'
	for i := 0; i < len(runes); i++ {
		c := runes[i]
		n, e := strconv.Atoi(string(c))
		if e == nil {
			res = do(res, n, op)
			continue
		}
		if c == '+' || c == '*' {
			op = c
		}
		if c == '(' {
			closing := findMatchingBracketPos(runes, i)
			insideExpr := string(runes[i+1 : closing])
			res = do(res, calc(insideExpr), op)
			i = closing
		}
	}
	return res
}

func main() {
	s := 0
	data, _ := ioutil.ReadFile("input.txt")
	text := strings.TrimSpace(string(data))
	for _, line := range strings.Split(text, "\n") {
		s += calc(line)
	}
	fmt.Println(s)
}
