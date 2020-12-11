package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func getData() [][]byte {
	data, _ := ioutil.ReadFile("input.txt")
	text := strings.TrimSpace(string(data))
	seats := [][]byte{}
	for _, line := range strings.Split(text, "\n") {
		seats = append(seats, []byte(line))
	}
	return seats
}

func BenchmarkOccupied(b *testing.B) {
	seats := getData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		occupiedSeatsWhenStabilized(seats, 5, -1)
	}
}

func BenchmarkOccupiedNoalloc(b *testing.B) {
	seats := getData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		occupiedSeatsWhenStabilizedNoalloc(seats, 5, -1)
	}
}
