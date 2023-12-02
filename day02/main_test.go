package main

import "testing"

func TestPart1(t *testing.T) {
	got := part1([]string{""})
	want := 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
