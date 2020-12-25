package main

import (
	"testing"
)

func assumeEqual(t *testing.T, got, expected int) {
	if got != expected {
		t.Errorf("Got %d, expected %d", got, expected)
	}
}

func TestCalcAst(t *testing.T) {
	assumeEqual(t, calcAst("1 + 2 * 3 + 4 * 5 + 6"), 231)
	assumeEqual(t, calcAst("1 + (2 * 3) + (4 * (5 + 6))"), 51)
	assumeEqual(t, calcAst("2 * 3 + (4 * 5)"), 46)
	assumeEqual(t, calcAst("5 + (8 * 3 + 9 + 3 * 4 * 3)"), 1445)
	assumeEqual(t, calcAst("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"), 669060)
	assumeEqual(t, calcAst("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"), 23340)
}
