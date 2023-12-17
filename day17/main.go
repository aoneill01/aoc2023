package main

import (
	"aoc2023/util"
	"fmt"
	"strconv"

	"github.com/albertorestifo/dijkstra"
)

const format = "{%d, %d}~%s"

var hDirs = []struct {
	row int
	col int
}{{0, 1}, {0, -1}}

var vDirs = []struct {
	row int
	col int
}{{1, 0}, {-1, 0}}

func main() {
	input := util.ReadDay(17)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	g := generateGraph(input)

	_, cost, _ := g.Path("{0, 0}", fmt.Sprintf("{%d, %d}", len(input)-1, len(input[0])-1))

	return cost
}

func part2(input []string) int {
	g := generateGraph2(input)

	_, cost, _ := g.Path("{0, 0}", fmt.Sprintf("{%d, %d}", len(input)-1, len(input[0])-1))

	return cost
}

func generateGraph(input []string) dijkstra.Graph {
	var g = dijkstra.Graph{}

	for row := range input {
		for col := range input[0] {
			addToGraph(row, col, input, g)
		}
	}

	return g
}

func addToGraph(row, col int, input []string, g dijkstra.Graph) {
	if row == len(input)-1 && col == len(input[0])-1 {
		g[fmt.Sprintf("{%d, %d}", len(input)-1, len(input[0])-1)] = map[string]int{}
		return
	}

	m := map[string]int{}
	for _, notDir := range []string{"h", "v"} {
		if row != 0 || col != 0 {
			m = map[string]int{}
		}

		dirs := hDirs
		if notDir == "h" {
			dirs = vDirs
		}
		for _, dir := range dirs {
			sum := 0
			for j := 1; j <= 3; j++ {
				r := row + j*dir.row
				c := col + j*dir.col
				if r < 0 || r >= len(input) || c < 0 || c >= len(input[0]) {
					break
				}
				v, _ := strconv.Atoi(string(input[r][c]))
				sum += v
				tmp := "h"
				if notDir == "h" {
					tmp = "v"
				}
				if r == len(input)-1 && c == len(input[0])-1 {
					m[fmt.Sprintf("{%d, %d}", r, c)] = sum
				} else {
					m[fmt.Sprintf(format, r, c, tmp)] = sum
				}
			}
		}
		if row == 0 && col == 0 {
			g["{0, 0}"] = m
		} else {
			g[fmt.Sprintf(format, row, col, notDir)] = m
		}
	}
}

func generateGraph2(input []string) dijkstra.Graph {
	var g = dijkstra.Graph{}

	for row := range input {
		for col := range input[0] {
			addToGraph2(row, col, input, g)
		}
	}

	return g
}

func addToGraph2(row, col int, input []string, g dijkstra.Graph) {
	if row == len(input)-1 && col == len(input[0])-1 {
		g[fmt.Sprintf("{%d, %d}", len(input)-1, len(input[0])-1)] = map[string]int{}
		return
	}

	m := map[string]int{}
	for _, notDir := range []string{"h", "v"} {
		if row != 0 || col != 0 {
			m = map[string]int{}
		}

		dirs := hDirs
		if notDir == "h" {
			dirs = vDirs
		}
		for _, dir := range dirs {
			sum := 0
			for j := 1; j < 4; j++ {
				r := row + j*dir.row
				c := col + j*dir.col
				if r < 0 || r >= len(input) || c < 0 || c >= len(input[0]) {
					break
				}
				v, _ := strconv.Atoi(string(input[r][c]))
				sum += v
			}
			for j := 4; j <= 10; j++ {
				r := row + j*dir.row
				c := col + j*dir.col
				if r < 0 || r >= len(input) || c < 0 || c >= len(input[0]) {
					break
				}
				v, _ := strconv.Atoi(string(input[r][c]))
				sum += v
				tmp := "h"
				if notDir == "h" {
					tmp = "v"
				}
				if r == len(input)-1 && c == len(input[0])-1 {
					m[fmt.Sprintf("{%d, %d}", r, c)] = sum
				} else {
					m[fmt.Sprintf(format, r, c, tmp)] = sum
				}
			}
		}
		if row == 0 && col == 0 {
			g["{0, 0}"] = m
		} else {
			g[fmt.Sprintf(format, row, col, notDir)] = m
		}
	}
}
