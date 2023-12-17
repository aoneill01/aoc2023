package main

import (
	"aoc2023/util"
	"fmt"
	"strconv"

	"github.com/albertorestifo/dijkstra"
)

const format = "{%d, %d}~%+v"

var dirs = []struct {
	row int
	col int
}{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func main() {
	input := util.ReadDay(17)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	g := generateGraph(input)

	path, cost, _ := g.Path("{0, 0}~{row:-1 col:0}", fmt.Sprintf(format, len(input)-1, len(input[0])-1, dirs[1])) // skipping error handling
	fmt.Println(path)
	return cost
}

func part2(input []string) int {
	return 0
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

	for notIndex := range dirs {
		m := map[string]int{}
		for i, dir := range dirs {
			if i == notIndex {
				continue
			}
			sum := 0
			for j := 1; j <= 3; j++ {
				r := row + j*dir.row
				c := col + j*dir.col
				if r < 0 || r >= len(input) || c < 0 || c >= len(input[0]) {
					break
				}
				v, _ := strconv.Atoi(string(input[r][c]))
				sum += v
				m[fmt.Sprintf(format, r, c, dir)] = sum
			}
		}
		g[fmt.Sprintf(format, row, col, dirs[notIndex])] = m
	}
}
