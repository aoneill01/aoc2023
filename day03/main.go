package main

import (
	"aoc2023/util"
	"fmt"
	"math"
	"strconv"
)

func main() {
	input := util.ReadDay(3)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

const dot = byte('.')
const gear = byte('*')

func isSymbol(b byte) bool {
	_, err := strconv.Atoi(string(b))
	return err != nil && b != dot
}

func isInBounds(input []string, row, col int) bool {
	return row >= 0 && col >= 0 && row < len(input) && col < len(input[row])
}

func hasAjacentSymbol(input []string, row, col int) bool {
	for r := row - 1; r <= row+1; r++ {
		for c := col - 1; c <= col+1; c++ {
			if isInBounds(input, r, c) && isSymbol(input[r][c]) {
				return true
			}
		}
	}
	return false
}

func part1(input []string) int {
	var result int

	for row, line := range input {
		num := 0
		seenSymbol := false
		for col := 0; col < len(line); col++ {
			c := line[col]
			if i, err := strconv.Atoi(string(c)); err == nil {
				num = 10*num + i
				if hasAjacentSymbol(input, row, col) {
					seenSymbol = true
				}
			} else {
				if num != 0 && seenSymbol {
					result += num
				}
				num = 0
				seenSymbol = false
			}
		}
		if num != 0 && seenSymbol {
			result += num
		}
	}
	return result
}

func findTouching(input []string, row, gearCol int) []int {
	result := []int{}

	if row < 0 || row >= len(input) {
		return result
	}

	line := input[row]
	num := 0
	seenGear := false
	for col := 0; col < len(line); col++ {
		c := line[col]
		if i, err := strconv.Atoi(string(c)); err == nil {
			num = 10*num + i
			if math.Abs(float64(col-gearCol)) <= 1 {
				seenGear = true
			}
		} else {
			if num != 0 && seenGear {
				result = append(result, num)
			}
			num = 0
			seenGear = false
		}
	}
	if num != 0 && seenGear {
		result = append(result, num)
	}

	return result
}

func part2(input []string) int {
	var result int

	for row, line := range input {
		for col := 0; col < len(line); col++ {
			c := line[col]
			if c == gear {
				touching := findTouching(input, row-1, col)
				touching = append(touching, findTouching(input, row, col)...)
				touching = append(touching, findTouching(input, row+1, col)...)
				if len(touching) == 2 {
					result += touching[0] * touching[1]
				}
			}
		}
	}
	return result
}
