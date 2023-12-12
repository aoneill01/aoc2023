package main

import (
	"strings"
	"testing"
)

const sample1 = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestPart1(t *testing.T) {
	got := part1(strings.Split(sample1, "\n"))
	want := 374

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestWithExpansion(t *testing.T) {
	got := withExpansion(strings.Split(sample1, "\n"), 9)
	want := 1030

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
