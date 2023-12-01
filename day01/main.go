package main

import (
	"aoc2023/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadDay(1)

	fmt.Println(Part1(input))
	fmt.Println(Part2(input))
}

func Part1(input []string) int {
	var result int

	for _, line := range input {
		result += CalibrationValue(line)
	}

	return result
}

func Part2(input []string) int {
	var result int

	for _, line := range input {
		result += CalibrationValue2(line)
	}

	return result
}

func CalibrationValue(line string) int {
	var value int

	for i := 0; i < len(line); i++ {
		if digit, err := strconv.Atoi(string(line[i])); err == nil {
			value = 10 * digit
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		if digit, err := strconv.Atoi(string(line[i])); err == nil {
			value += digit
			break
		}
	}

	return value
}

func CalibrationValue2(line string) int {
	return 10*firstDigit(line, false) + firstDigit(line, true)
}

type Num = struct {
	name  string
	value int
}

func firstDigit(line string, reversed bool) int {
	numbers := []*Num{
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
	}

	if reversed {
		for _, number := range numbers {
			number.name = reverse(number.name)
		}
		line = reverse(line)
	}

	for i := 0; i < len(line); i++ {
		for _, number := range numbers {
			if strings.HasPrefix(line[i:], number.name) {
				return number.value
			}
		}
		if digit, err := strconv.Atoi(string(line[i])); err == nil {
			return digit
		}
	}

	panic("Expected digit")
}

func reverse(s string) string {
	var byte strings.Builder
	byte.Grow(len(s))
	for i := len(s) - 1; i >= 0; i-- {
		byte.WriteByte(s[i])
	}
	return byte.String()
}
