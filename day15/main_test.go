package main

import (
	"testing"
)

const sample1 = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

func TestPart1(t *testing.T) {
	got := part1(sample1)
	want := 1320

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(sample1)
	want := 145

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
