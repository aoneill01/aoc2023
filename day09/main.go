package main

import (
	"aoc2023/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadDay(9)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	result := 0

	for _, line := range input {
		v := nextValue(parseNumbers(line))
		result += v
	}

	return result
}

func part2(input []string) int {
	result := 0

	for _, line := range input {
		v := previousValue(parseNumbers(line))
		result += v
	}

	return result
}

func parseNumbers(list string) []int {
	result := []int{}

	for _, n := range strings.Split(list, " ") {
		if n == "" {
			continue
		}
		v, _ := strconv.Atoi(n)
		result = append(result, v)
	}

	return result
}

func nextValue(values []int) int {
	diffs := make([]int, 0, len(values)-1)
	allZeros := true

	for i := 0; i < len(values)-1; i++ {
		diff := values[i+1] - values[i]
		if diff != 0 {
			allZeros = false
		}
		diffs = append(diffs, diff)
	}

	if allZeros {
		return values[len(values)-1]
	}

	return values[len(values)-1] + nextValue(diffs)
}

func previousValue(values []int) int {
	diffs := make([]int, 0, len(values)-1)
	allZeros := true

	for i := 0; i < len(values)-1; i++ {
		diff := values[i+1] - values[i]
		if diff != 0 {
			allZeros = false
		}
		diffs = append(diffs, diff)
	}

	if allZeros {
		return values[0]
	}

	return values[0] - previousValue(diffs)
}
