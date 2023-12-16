package main

import (
	"strings"
	"testing"
)

const sample1 = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

func TestPart1(t *testing.T) {
	got := part1(strings.Split(sample1, "\n"))
	want := 46

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(strings.Split(sample1, "\n"))
	want := 51

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
