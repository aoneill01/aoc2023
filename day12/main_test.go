package main

import (
	"strings"
	"testing"
)

const sample1 = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func TestPart1(t *testing.T) {
	got := part1(strings.Split(sample1, "\n"))
	want := 21

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(strings.Split(sample1, "\n"))
	want := 525152

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
