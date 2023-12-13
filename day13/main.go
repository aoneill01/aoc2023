package main

import (
	"aoc2023/util"
	"fmt"
)

func main() {
	input := util.ReadDay(13)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	patterns := parseInput(input)

	result := 0

	for _, pattern := range patterns {
		result += summarize(pattern)
	}

	return result
}

func part2(input []string) int {
	patterns := parseInput(input)

	result := 0

	for _, pattern := range patterns {
		result += summarize2(pattern)
	}

	return result
}

func parseInput(input []string) [][]string {
	result := [][]string{}

	start := 0
	for i, val := range input {
		if val == "" {
			result = append(result, input[start:i])
			start = i + 1
		}
	}
	result = append(result, input[start:])

	return result
}

func summarize(pattern []string) int {
testLabel:
	for testCol := 0; testCol < len(pattern[0])-1; testCol++ {
		for reflect := 0; reflect < len(pattern[0]); reflect++ {
			c1 := testCol - reflect
			c2 := testCol + reflect + 1
			if c1 < 0 || c1 >= len(pattern[0]) || c2 < 0 || c2 >= len(pattern[0]) {
				continue
			}
			for r := range pattern {
				if pattern[r][c1] != pattern[r][c2] {
					continue testLabel
				}
			}
		}
		return testCol + 1
	}

testLabel2:
	for testRow := 0; testRow < len(pattern)-1; testRow++ {
		for reflect := 0; reflect < len(pattern); reflect++ {
			r1 := testRow - reflect
			r2 := testRow + reflect + 1
			if r1 < 0 || r1 >= len(pattern) || r2 < 0 || r2 >= len(pattern) {
				continue
			}
			for c := range pattern[0] {
				if pattern[r1][c] != pattern[r2][c] {
					continue testLabel2
				}
			}
		}
		return 100 * (testRow + 1)
	}

	return 0
}

func summarize2(pattern []string) int {
testLabel:
	for testCol := 0; testCol < len(pattern[0])-1; testCol++ {
		smudgeCount := 0
		for reflect := 0; reflect < len(pattern[0]); reflect++ {
			c1 := testCol - reflect
			c2 := testCol + reflect + 1
			if c1 < 0 || c1 >= len(pattern[0]) || c2 < 0 || c2 >= len(pattern[0]) {
				continue
			}
			for r := range pattern {
				if pattern[r][c1] != pattern[r][c2] {
					smudgeCount++
					if smudgeCount > 1 {
						continue testLabel
					}
				}
			}
		}
		if smudgeCount == 1 {
			return testCol + 1
		}
	}

testLabel2:
	for testRow := 0; testRow < len(pattern)-1; testRow++ {
		smudgeCount := 0
		for reflect := 0; reflect < len(pattern); reflect++ {
			r1 := testRow - reflect
			r2 := testRow + reflect + 1
			if r1 < 0 || r1 >= len(pattern) || r2 < 0 || r2 >= len(pattern) {
				continue
			}
			for c := range pattern[0] {
				if pattern[r1][c] != pattern[r2][c] {
					smudgeCount++
					if smudgeCount > 1 {
						continue testLabel2
					}
				}
			}
		}
		if smudgeCount == 1 {
			return 100 * (testRow + 1)
		}
	}

	return 0
}
