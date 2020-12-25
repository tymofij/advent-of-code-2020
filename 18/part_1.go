package main

import "strconv"

func simpleCalc(s string) int {
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
			res = do(res, simpleCalc(insideExpr), op)
			i = closing
		}
	}
	return res
}
