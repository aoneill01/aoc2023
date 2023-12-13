package main

import (
	"aoc2023/util"
	"fmt"
	"strconv"
	"strings"
)

type row struct {
	condition string
	springs   []int
}

type key struct {
	i int
	j int
}

func main() {
	input := util.ReadDay(12)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	rows := parseInput(input)

	count := 0

	for _, row := range rows {
		cache := make(map[key]int)
		count += possibleArrangements(0, 0, row, cache)
	}

	return count
}

func part2(input []string) int {
	originalRows := parseInput(input)
	rows := []row{}
	for _, row := range originalRows {
		rows = append(rows, expandRow(row))
	}

	count := 0

	for _, row := range rows {
		cache := make(map[key]int)
		count += possibleArrangements(0, 0, row, cache)
	}

	return count
}

func parseInput(input []string) []row {
	results := []row{}

	for _, line := range input {
		condition, after, _ := strings.Cut(line, " ")

		springs := []int{}

		for _, n := range strings.Split(after, ",") {
			v, _ := strconv.Atoi(n)
			springs = append(springs, v)
		}

		results = append(results, row{condition, springs})
	}

	return results
}

func possibleArrangements(conditionIndex int, springIndex int, r row, cache map[key]int) int {
	// Are there enough remaining to have a possible solution?
	minLength := len(r.springs) - springIndex - 1
	if minLength < 0 {
		minLength = 0
	}
	for i := springIndex; i < len(r.springs); i++ {
		minLength += r.springs[i]
	}
	if len(r.condition)-conditionIndex < minLength {
		cache[key{conditionIndex, springIndex}] = 0
		return 0
	}

	// No remaining springs to process
	if springIndex == len(r.springs) {
		for i := conditionIndex; i < len(r.condition); i++ {
			if r.condition[i] == '#' {
				cache[key{conditionIndex, springIndex}] = 0
				return 0
			}
		}
		cache[key{conditionIndex, springIndex}] = 1
		return 1
	}

	count := 0
	springRun := r.springs[springIndex]

outer:
	for startIndex := conditionIndex; startIndex < len(r.condition); startIndex++ {
		for delta := 0; delta < springRun; delta++ {
			if startIndex+delta >= len(r.condition) {
				break outer
			}
			if r.condition[startIndex+delta] == '.' {
				if r.condition[startIndex] == '#' {
					break outer
				}
				continue outer
			}
		}
		if startIndex+springRun == len(r.condition) {
			k := key{startIndex + springRun, springIndex + 1}
			var tmp int
			if v, ok := cache[k]; ok {
				tmp = v
			} else {
				tmp = possibleArrangements(startIndex+springRun, springIndex+1, r, cache)
			}
			count += tmp
		} else if r.condition[startIndex+springRun] != '#' {
			k := key{startIndex + springRun + 1, springIndex + 1}
			var tmp int
			if v, ok := cache[k]; ok {
				tmp = v
			} else {
				tmp = possibleArrangements(startIndex+springRun+1, springIndex+1, r, cache)
			}
			count += tmp
		}

		if r.condition[startIndex] == '#' {
			break outer
		}
	}

	cache[key{conditionIndex, springIndex}] = count
	return count
}

func expandRow(r row) row {
	condition := strings.Join([]string{r.condition, r.condition, r.condition, r.condition, r.condition}, "?")
	springs := []int{}
	for i := 0; i < 5; i++ {
		for _, s := range r.springs {
			springs = append(springs, s)
		}
	}
	return row{condition, springs}
}
