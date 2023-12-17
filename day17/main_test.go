package main

import (
	"strings"
	"testing"
)

const sample1 = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

func TestPart1(t *testing.T) {
	got := part1(strings.Split(sample1, "\n"))
	want := 102

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := part2(strings.Split(sample1, "\n"))
	want := -1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

// map[
// 	{0, 0}~{row:-1 col:0}:map[
// 		{0, 1}~{row:0 col:1}:4 {0, 2}~{row:0 col:1}:5 {0, 3}~{row:0 col:1}:8 {1, 0}~{row:1 col:0}:3 {2, 0}~{row:1 col:0}:6 {3, 0}~{row:1 col:0}:9
// 	]
// 	{0, 0}~{row:0 col:-1}:map[
// 		{0, 1}~{row:0 col:1}:4 {0, 2}~{row:0 col:1}:5 {0, 3}~{row:0 col:1}:8 {1, 0}~{row:1 col:0}:3 {2, 0}~{row:1 col:0}:6 {3, 0}~{row:1 col:0}:9
// 	]
// 	{0, 0}~{row:0 col:1}:map[
// 		{1, 0}~{row:1 col:0}:3 {2, 0}~{row:1 col:0}:6 {3, 0}~{row:1 col:0}:9
// 	]
// 	{0, 0}~{row:1 col:0}:map[
// 		{0, 1}~{row:0 col:1}:4 {0, 2}~{row:0 col:1}:5 {0, 3}~{row:0 col:1}:8
// 	]
// ]

// {0, 0}~{row:-1 col:0}
// ~{row:-1 col:0}
// ~{row:0 col:-1}
