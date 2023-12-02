package main

import (
	"aoc2023/util"
	"fmt"
	"strconv"
	"strings"
)

type cubes struct {
	red   int
	green int
	blue  int
}

type game struct {
	id       int
	revealed []cubes
}

func main() {
	input := util.ReadDay(2)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func parseLine(line string) game {
	parts := strings.Split(line, ": ")
	idStr, _ := strings.CutPrefix(parts[0], "Game ")
	id, _ := strconv.Atoi(idStr)
	game := game{id, []cubes{}}

	parts = strings.Split(parts[1], "; ")
	for _, part := range parts {
		c := cubes{}
		draws := strings.Split(part, ", ")
		for _, draw := range draws {
			a := strings.Split(draw, " ")
			count, _ := strconv.Atoi(a[0])
			switch a[1] {
			case "red":
				c.red = count
			case "green":
				c.green = count
			case "blue":
				c.blue = count
			}
		}
		game.revealed = append(game.revealed, c)
	}

	return game
}

func isValidPart1(g game) bool {
	for _, c := range g.revealed {
		if c.red > 12 || c.green > 13 || c.blue > 14 {
			return false
		}
	}

	return true
}

func fewestPossible(g game) cubes {
	result := cubes{}

	for _, c := range g.revealed {
		if c.red > result.red {
			result.red = c.red
		}
		if c.green > result.green {
			result.green = c.green
		}
		if c.blue > result.blue {
			result.blue = c.blue
		}
	}

	return result
}

func part1(input []string) int {
	var result int

	for _, line := range input {
		g := parseLine(line)
		if isValidPart1(g) {
			result += g.id
		}
	}

	return result
}

func part2(input []string) int {
	var result int

	for _, line := range input {
		g := parseLine(line)
		fp := fewestPossible(g)
		result += fp.red * fp.green * fp.blue
	}

	return result
}
