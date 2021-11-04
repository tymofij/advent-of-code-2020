package main

import (
	"testing"

	"test18/a/mock_reader"

	"github.com/golang/mock/gomock"
)

func TestBracketPos(t *testing.T) {
	cases := []struct {
		s        string
		startPos int
		expected int
	}{{
		s:        "()",
		startPos: 0,
		expected: 1,
	},
		{
			s:        "x(4+5)",
			startPos: 1,
			expected: 5,
		},
		{
			s:        "x((1))",
			startPos: 1,
			expected: 5,
		},
	}
	for _, c := range cases {
		t.Run(c.s, func(t *testing.T) {
			result := findMatchingBracketPos([]rune(c.s), c.startPos)
			if result != c.expected {
				t.Errorf("Got %d, expected %d", result, c.expected)
			}
		})
	}
}

func TestSimpleCalc(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{
			input:    "2 * 3",
			expected: 6,
		},
		{
			input:    "2 + 3",
			expected: 5,
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

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			result := SimpleCalc(c.input)
			if result != c.expected {
				t.Errorf("Got %d, expected %d", result, c.expected)
			}
		})
	}

	// subtests for input having 2 can be run with `go test -v -run "TestSimpleCalc/2`
}

type StubReader struct{}

func (r StubReader) ReadLines() []string {
	return []string{
		"2+3",
		"2*3",
	}
}

func TestCalcSum(t *testing.T) {
	stubreader := StubReader{}
	calculator := NewCalcLinesSum(stubreader)
	result := calculator.calculate()
	expected := 11
	if result != expected {
		t.Errorf("Got %d, expected %d", result, expected)
	}
}

// mockgen -source=./reader/types.go -destination=./mock_reader/reader.go

func TestCalcSumMocked(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockReader := mock_reader.NewMockLineReader(mockCtrl)
	mockReader.EXPECT().ReadLines().Return([]string{"2+2"}).Times(1)
	calculator := NewCalcLinesSum(mockReader)
	result := calculator.calculate()
	expected := 4
	if result != expected {
		t.Errorf("Got %d, expected %d", result, expected)
	}
}
