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

func Part1(input []string) (sum int) {
	for _, line := range input {
		sum += CalibrationValue(line)
	}

	return
}

func Part2(input []string) (sum int) {
	for _, line := range input {
		sum += CalibrationValue2(line)
	}

	return
}

func CalibrationValue(line string) (result int) {
	for i := 0; i < len(line); i++ {
		if digit, err := strconv.Atoi(string(line[i])); err == nil {
			result = 10 * digit
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if digit, err := strconv.Atoi(string(line[i])); err == nil {
			result += digit
			break
		}
	}

	return
}

func CalibrationValue2(line string) int {
	numbers := toNumbers(line)

	return 10*numbers[0] + numbers[len(numbers)-1]
}

type Substitution = struct {
	name  string
	value int
}

var numbers = []Substitution{
	{"one", 1},
	{"two", 2},
	{"three", 3},
	{"four", 4},
	{"five", 5},
	{"six", 6},
	{"seven", 7},
	{"eight", 8},
	{"nine", 9},
	{"1", 1},
	{"2", 2},
	{"3", 3},
	{"4", 4},
	{"5", 5},
	{"6", 6},
	{"7", 7},
	{"8", 8},
	{"9", 9},
	{"0", 0},
}

func toNumbers(line string) []int {
	result := []int{}

	for i := 0; i < len(line); i++ {
		for _, number := range numbers {
			if strings.HasPrefix(line[i:], number.name) {
				result = append(result, number.value)
			}
		}
	}

	return result
}
