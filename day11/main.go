package main

import (
	"aoc2023/util"
	"fmt"
)

type star struct {
	x int
	y int
}

func main() {
	input := util.ReadDay(11)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	return withExpansion(input, 1)
}

func part2(input []string) int {
	return withExpansion(input, 1000000-1)
}

func withExpansion(input []string, expansion int) int {
	stars := parseInput(input, expansion)

	result := 0

	for i := 0; i < len(stars)-1; i++ {
		for j := i + 1; j < len(stars); j++ {
			result += starDistance(stars[i], stars[j])
		}
	}

	return result
}

func parseInput(input []string, expansion int) []star {
	yDeltas := make([]int, len(input))
	xDeltas := make([]int, len(input[0]))

o1:
	for y := range input {
		for x := range input[y] {
			if input[y][x] == '#' {
				if y == 0 {
					yDeltas[y] = 0
				} else {
					yDeltas[y] = yDeltas[y-1]
				}
				continue o1
			}
		}
		if y == 0 {
			yDeltas[y] = expansion
		} else {
			yDeltas[y] = yDeltas[y-1] + expansion
		}
	}

o2:
	for x := range input[0] {
		for y := range input {
			if input[y][x] == '#' {
				if x == 0 {
					xDeltas[x] = 0
				} else {
					xDeltas[x] = xDeltas[x-1]
				}
				continue o2
			}
		}
		if x == 0 {
			xDeltas[x] = expansion
		} else {
			xDeltas[x] = xDeltas[x-1] + expansion
		}
	}

	stars := []star{}

	for y := range input {
		for x := range input[y] {
			if input[y][x] == '#' {
				stars = append(stars, star{x + xDeltas[x], y + yDeltas[y]})
			}
		}
	}

	return stars
}

func starDistance(s1, s2 star) int {
	return abs(s1.x-s2.x) + abs(s1.y-s2.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
