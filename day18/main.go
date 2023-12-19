package main

import (
	"aoc2023/util"
	"fmt"
	"regexp"
	"strconv"
)

type instruction struct {
	direction string
	distance  int
}

type minMax struct {
	min int
	max int
}

type point struct {
	x int
	y int
}

func (mm minMax) update(val int) minMax {
	if val < mm.min {
		mm.min = val
	}
	if val > mm.max {
		mm.max = val
	}
	return mm
}

func (mm minMax) size() int {
	return mm.max - mm.min + 1
}

func (mm minMax) includes(val int) bool {
	return val >= mm.min && val <= mm.max
}

var re = regexp.MustCompile(`^(.) (\d+) \(#(.*)\)$`)
var deltas = map[string]struct {
	x int
	y int
}{
	"U": {0, -1},
	"R": {1, 0},
	"D": {0, 1},
	"L": {-1, 0},
}

func main() {
	input := util.ReadDay(18)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	rows := make(map[int]minMax)
	instructions := parseInput(input)
	var x, y, minX, minY, maxX, maxY int

	rows[0] = minMax{0, 0}

	for _, instruction := range instructions {
		delta := deltas[instruction.direction]
		for i := 0; i < instruction.distance; i++ {
			x += delta.x
			y += delta.y
			if x < minX {
				minX = x
			}
			if y < minY {
				minY = y
			}
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	grid := make([][]byte, maxY-minY+3)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]byte, maxX-minX+3)
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = byte('.')
		}
	}

	grid[0-minY+1][0-minX+1] = byte('#')
	for _, instruction := range instructions {
		delta := deltas[instruction.direction]
		for i := 0; i < instruction.distance; i++ {
			x += delta.x
			y += delta.y
			grid[y-minY+1][x-minX+1] = byte('#')
		}
	}

	fill(0, 0, grid)

	result := 0

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] != ' ' {
				result++
			}
		}
	}

	return result
}

func part2(input []string) int {
	instructions := parseInput2(input)
	var x, y int
	points1 := make([]point, 0, len(instructions))
	points2 := make([]point, 0, len(instructions))

	for i, instr := range instructions {
		delta := deltas[instr.direction]
		x += instr.distance * delta.x
		y += instr.distance * delta.y

		var deltaX, deltaY int
		var nextInstr instruction
		if i == len(instructions)-1 {
			nextInstr = instructions[len(instructions)-1]
		} else {
			nextInstr = instructions[i+1]
		}
		if nextInstr.direction == "D" || instr.direction == "D" {
			deltaX = 1
		}
		if nextInstr.direction == "L" || instr.direction == "L" {
			deltaY = 1
		}

		points1 = append(points1, point{x + deltaX, y + deltaY})
		deltaX = 1 - deltaX
		deltaY = 1 - deltaY
		points2 = append(points2, point{x + deltaX, y + deltaY})
	}

	var result1, result2 int

	for i := 0; i < len(points1); i++ {
		a1 := points1[i]
		a2 := points2[i]
		var b1, b2 point
		if i < len(points1)-1 {
			b1 = points1[i+1]
			b2 = points2[i+1]
		} else {
			b1 = points1[0]
			b2 = points2[0]
		}

		result1 += a1.y * (a1.x - b1.x)
		result2 += a2.y * (a2.x - b2.x)
	}

	if result1 > result2 {
		return result1
	}
	return result2
}

func parseInput(input []string) []instruction {
	result := make([]instruction, 0, len(input))

	for _, line := range input {
		matches := re.FindStringSubmatch(line)
		distance, _ := strconv.Atoi(matches[2])
		result = append(result, instruction{matches[1], distance})
	}

	return result
}

func fill(x, y int, grid [][]byte) {
	if x < 0 || y < 0 || x >= len(grid[0]) || y >= len(grid) || grid[y][x] != '.' {
		return
	}

	grid[y][x] = byte(' ')
	fill(x-1, y, grid)
	fill(x+1, y, grid)
	fill(x, y-1, grid)
	fill(x, y+1, grid)
}

func parseInput2(input []string) []instruction {
	digitToDirection := map[string]string{
		"0": "R",
		"1": "D",
		"2": "L",
		"3": "U",
	}
	result := make([]instruction, 0, len(input))

	for _, line := range input {
		matches := re.FindStringSubmatch(line)
		distance, _ := strconv.ParseInt(matches[3][:5], 16, 64)
		result = append(result, instruction{digitToDirection[matches[3][5:]], int(distance)})
	}

	return result
}
