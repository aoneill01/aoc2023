package main

import (
	"strings"
	"testing"
)

const sample1 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

const sample2 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func TestPart1(t *testing.T) {
	got := part1(strings.Split(sample1, "\n"))
	want := 6

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(strings.Split(sample2, "\n"))
	want := 6

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
