package main

import (
	"aoc2023/util"
	"fmt"
	"regexp"
	"sort"
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

type verticalLine struct {
	x         int
	yRange    minMax
	direction string
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

func (mm minMax) onEdge(val int) bool {
	return val == mm.min || val == mm.max
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

	//	fmt.Print(instructions)
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

	for _, line := range grid {
		fmt.Println(string(line))
	}

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
	var x, y, minY int
	verticalLines := []verticalLine{}
	yStopsMap := map[int]bool{}

	for _, instruction := range instructions {
		delta := deltas[instruction.direction]
		origY := y
		x += instruction.distance * delta.x
		y += instruction.distance * delta.y

		if origY != y {
			verticalLines = append(verticalLines, verticalLine{x, minMax{y, y}.update(origY), instruction.direction})
			yStopsMap[y] = true
			yStopsMap[origY] = true
		}
		if y < minY {
			minY = y
		}
	}

	yStops := []int{}
	for yStop := range yStopsMap {
		yStops = append(yStops, yStop)
	}
	sort.Ints(yStops)
	fmt.Println(yStops)

	sort.Slice(verticalLines, func(i, j int) bool {
		return verticalLines[i].x < verticalLines[j].x
	})

	fmt.Println(x, y, verticalLines)

	result := 0

	for i := 0; i < len(yStops); i++ {
		w := width(yStops[i], verticalLines)
		result += w

		fmt.Println("A", yStops[i], w)

		if i < len(yStops)-1 && yStops[i+1] != yStops[i]+1 {
			w = width(yStops[i]+1, verticalLines)
			h := yStops[i+1] - yStops[i] - 1
			fmt.Println("B", yStops[i]+1, w, h, w*h)
			result += w * h
		}
	}

	return result

	//y = minY
	// for {
	// 	intersecting := []verticalLine{}
	// 	for _, l := range verticalLines {
	// 		if l.yRange.includes(y) {
	// 			intersecting = append(intersecting, l)
	// 		}
	// 	}

	// 	if len(intersecting) == 0 {
	// 		return result
	// 	}
	// 	fmt.Println(len(intersecting))

	// 	nextY := intersecting[0].yRange.max
	// 	for _, l := range intersecting {
	// 		if l.yRange.max < nextY {
	// 			nextY = l.yRange.max
	// 		}
	// 	}
	// 	if nextY == y {
	// 		nextY++
	// 	}
	// 	firstDirection := intersecting[0].direction
	// 	filtered := []verticalLine{}
	// 	for i, l := range intersecting {
	// 		if l.direction == firstDirection && (i > 0 && intersecting[i-1].direction == firstDirection) {
	// 			continue
	// 		}
	// 		if l.direction != firstDirection && (i < len(intersecting)-1 && intersecting[i+1].direction != firstDirection) {
	// 			continue
	// 		}
	// 		filtered = append(filtered, l)
	// 	}
	// 	if len(filtered)%2 != 0 {
	// 		fmt.Println("BAD!")
	// 	}
	// 	width := 0
	// 	for i := 0; i < len(filtered); i += 2 {
	// 		width += filtered[i+1].x - filtered[i].x + 1
	// 	}
	// 	result += width * (nextY - y)
	// 	fmt.Println(y, nextY, width, nextY-y, result, intersecting)

	// 	y = nextY

	// }

}

func width(y int, verticalLines []verticalLine) int {
	intersecting := []verticalLine{}
	for _, l := range verticalLines {
		if l.yRange.includes(y) {
			intersecting = append(intersecting, l)
		}
	}

	firstDirection := intersecting[0].direction
	filtered := []verticalLine{}
	for i, l := range intersecting {
		if l.direction == firstDirection && (i > 0 && intersecting[i-1].direction == firstDirection) {
			continue
		}
		if l.direction != firstDirection && (i < len(intersecting)-1 && intersecting[i+1].direction != firstDirection) {
			continue
		}
		filtered = append(filtered, l)
	}
	if len(filtered)%2 != 0 {
		fmt.Println("BAD!")
	}
	width := 0
	for i := 0; i < len(filtered); i += 2 {
		width += filtered[i+1].x - filtered[i].x + 1
	}
	for i := 1; i < len(filtered)-1; i += 2 {
		if y == filtered[i].yRange.min && filtered[i].yRange.min == filtered[i+1].yRange.min {
			fmt.Println("!!!", y)
			width += filtered[i+1].x - filtered[i].x - 1
		}
		// if y == filtered[i].yRange.max && filtered[i].yRange.max == filtered[i+1].yRange.max {
		// 	fmt.Println("???", y)
		// 	width += filtered[i+1].x - filtered[i].x - 1
		// }
	}
	return width
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
