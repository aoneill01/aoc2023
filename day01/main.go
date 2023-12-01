package main

import (
	"aoc2023/util"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	input := util.ReadDay(1)

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}

func sortedElves(input []string) []int {
	var elves []int
	var currentSum int

	for _, line := range input {
		if line == "" {
			elves = append(elves, currentSum)
			currentSum = 0
			continue
		}

		calories, _ := strconv.Atoi(line)
		currentSum += calories
	}

	elves = append(elves, currentSum)

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	return elves
}

func Part1(input []string) int {
	return sortedElves(input)[0]
}

func Part2(input []string) int {
	elves := sortedElves(input)

	return elves[0] + elves[1] + elves[2]
}
