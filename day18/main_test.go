package main

import (
	"strings"
	"testing"
)

const sample1 = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

func TestPart1(t *testing.T) {
	got := part1(strings.Split(sample1, "\n"))
	want := 62

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(strings.Split(sample1, "\n"))
	want := 952408144115
	// want := 62

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

/* U     U       D      D       U       D       D
   0     5411    461937 497056  609066  818608  1186328

   0--------------0
   .              .
   .              56407------------------56407
   .                     356353--356353  .
   500254-500254         .       .       .
          .              .       .       919647--919647
		  .              .       .               .
		  .              .       .               .
		  1186328--------1186328 1186328---------1186328













*/
