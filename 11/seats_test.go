package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func getData(name string) [][]byte {
	data, _ := ioutil.ReadFile(name)
	text := strings.TrimSpace(string(data))
	seats := [][]byte{}
	for _, line := range strings.Split(text, "\n") {
		seats = append(seats, []byte(line))
	}
	return seats
}

func TestNearby(t *testing.T) {
	seats := getData("i.txt")
	got := occupiedSeatsWhenStabilized(seats, 4, 1)
	if got != 37 {
		t.Errorf("got %d for occupiedSeatsWhenStabilized", got)
	}
}

func TestNearbyNoalloc(t *testing.T) {
	seats := getData("i.txt")
	got := occupiedSeatsWhenStabilizedNoalloc(seats, 4, 1)
	if got != 37 {
		t.Errorf("got %d for occupiedSeatsWhenStabilizedNoalloc", got)
	}
}

func TestVisible(t *testing.T) {
	seats := getData("i.txt")
	got := occupiedSeatsWhenStabilized(seats, 5, -1)
	if got != 26 {
		t.Errorf("got %d for occupiedSeatsWhenStabilized", got)
	}
}

func TestVisibleNoalloc(t *testing.T) {
	seats := getData("i.txt")
	got := occupiedSeatsWhenStabilizedNoalloc(seats, 5, -1)
	if got != 26 {
		t.Errorf("got %d for occupiedSeatsWhenStabilizedNoalloc", got)
	}
}

func BenchmarkOccupied(b *testing.B) {
	seats := getData("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		occupiedSeatsWhenStabilized(seats, 5, -1)
	}
}

func BenchmarkOccupiedNoalloc(b *testing.B) {
	seats := getData("input.txt")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		occupiedSeatsWhenStabilizedNoalloc(seats, 5, -1)
	}
}
