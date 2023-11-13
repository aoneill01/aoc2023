package main

import (
	"aoc2023/util"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	input := util.ReadDay(1)

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

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	fmt.Println(elves[0])
	fmt.Println(elves[0] + elves[1] + elves[2])
}
