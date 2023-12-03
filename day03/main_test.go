package main

import (
	"strings"
	"testing"
)

const sample1 = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestPart1(t *testing.T) {
	got := part1(strings.Split(sample1, "\n"))
	want := 4361

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(strings.Split(sample1, "\n"))
	want := 467835

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
