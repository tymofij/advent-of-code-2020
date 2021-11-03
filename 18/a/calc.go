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

// calculate the arithmentic expression,
// but addition and multiplication have SAME level of precedence
// except when there are parentheses. Parentheses still matter.
func SimpleCalc(s string) int {
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
			res = do(res, SimpleCalc(insideExpr), op)
			i = closing
		}
	}
	return res
}

func main() {
	data, _ := ioutil.ReadFile("../input.txt")
	text := strings.TrimSpace(string(data))
	lines := strings.Split(text, "\n")

	s := 0
	for _, line := range lines {
		s += SimpleCalc(line)
	}
	fmt.Println(s)
}
