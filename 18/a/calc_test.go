package main

import (
	"testing"
)

func assumeEqual(t *testing.T, got, expected int) {
	if got != expected {
		t.Errorf("Got %d, expected %d", got, expected)
	}
}
func TestBracketPos(t *testing.T) {
	assumeEqual(t, findMatchingBracketPos([]rune("()"), 0), 1)
	assumeEqual(t, findMatchingBracketPos([]rune("x(4+5)"), 1), 5)
	assumeEqual(t, findMatchingBracketPos([]rune("x((1))"), 1), 5)
}

func TestSimpleCalc(t *testing.T) {
	assumeEqual(t, simpleCalc("2 * 3"), 6)
	assumeEqual(t, simpleCalc("2 + 3"), 5)
	assumeEqual(t, simpleCalc("2 + (1 + 2)"), 5)
	assumeEqual(t, simpleCalc("1 + 2 * 3 + 4 * 5 + 6"), 71)
	assumeEqual(t, simpleCalc("2 * 3 + (4 * 5)"), 26)
	assumeEqual(t, simpleCalc("5 + (8 * 3 + 9 + 3 * 4 * 3)"), 437)
	assumeEqual(t, simpleCalc("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"), 12240)
	assumeEqual(t, simpleCalc("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"), 13632)
}
