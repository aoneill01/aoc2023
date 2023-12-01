package main

import (
	"strings"
	"testing"
)

const sampleInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestPart1(t *testing.T) {
	got := Part1(strings.Split(sampleInput, "\n"))
	want := 24000

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(strings.Split(sampleInput, "\n"))
	want := 45000

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
