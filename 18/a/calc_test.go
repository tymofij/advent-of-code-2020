package main

import (
	"testing"
)

func TestSimpleCalc(t *testing.T) {
	var result int

	result = SimpleCalc("2 * 3")
	if result != 6 {
		t.Errorf("Got %d, expected %d", result, 6)
	}

	result = SimpleCalc("1 + 2 * 3 + 4 * 5 + 6")
	if result != 71 {
		t.Errorf("Got %d, expected %d", result, 71)
	}

	// "2 * 3 + (4 * 5)" == 26
	// "5 + (8 * 3 + 9 + 3 * 4 * 3)" == 437
	// "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))" == 12240
	// "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2" == 13632
}
