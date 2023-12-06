package main

import (
	"strings"
	"testing"
)

const sample1 = `Time:      7  15   30
Distance:  9  40  200`

func TestPart1(t *testing.T) {
	got := part1(strings.Split(sample1, "\n"))
	want := 288

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(strings.Split(sample1, "\n"))
	want := 71503

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
