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
