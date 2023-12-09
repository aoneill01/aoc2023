package main

import (
	"strings"
	"testing"
)

const sample1 = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestPart1(t *testing.T) {
	got := part1(strings.Split(sample1, "\n"))
	want := 114

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(strings.Split(sample1, "\n"))
	want := 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
