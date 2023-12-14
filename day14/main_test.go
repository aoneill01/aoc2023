package main

import (
	"strings"
	"testing"
)

const sample1 = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

func TestPart1(t *testing.T) {
	got := part1(strings.Split(sample1, "\n"))
	want := 136

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(strings.Split(sample1, "\n"))
	want := 64

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
