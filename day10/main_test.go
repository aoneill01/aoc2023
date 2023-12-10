package main

import (
	"strings"
	"testing"
)

const sample1 = `-L|F7
7S-7|
L|7||
-L-J|
L|-JF`

const sample2 = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

const sample3 = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

const sample4 = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

func gotWantHelper(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart1(t *testing.T) {
	t.Run("sample input 1", func(t *testing.T) {
		got := part1(strings.Split(sample1, "\n"))
		gotWantHelper(t, got, 4)
	})

	t.Run("sample input 2", func(t *testing.T) {
		got := part1(strings.Split(sample2, "\n"))
		gotWantHelper(t, got, 8)
	})
}

func TestPart2(t *testing.T) {
	t.Run("sample input 3", func(t *testing.T) {
		got := part2(strings.Split(sample3, "\n"))
		gotWantHelper(t, got, 4)
	})

	t.Run("sample input 4", func(t *testing.T) {
		got := part2(strings.Split(sample4, "\n"))
		gotWantHelper(t, got, 10)
	})
}
