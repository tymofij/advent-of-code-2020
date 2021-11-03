package main

import (
	"testing"
)

type testCase struct {
	input    string
	expected int
}

func TestSimpleCalc(t *testing.T) {
	cases := []testCase{
		{
			input:    "2 * 3",
			expected: 6,
		},
		{
			input:    "1 + 2 * 3 + 4 * 5 + 6",
			expected: 71,
		},
		{
			input:    "2 * 3 + (4 * 5)",
			expected: 26,
		},
		{
			input:    "5 + (8 * 3 + 9 + 3 * 4 * 3)",
			expected: 437,
		},
		{
			input:    "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			expected: 12240,
		},
		{
			input:    "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			expected: 13632,
		},
	}

	var result int
	for _, c := range cases {
		result = SimpleCalc(c.input)
		if result != c.expected {
			t.Errorf("Got %d, expected %d", result, c.expected)
		}
	}

}
