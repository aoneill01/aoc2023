package main

import (
	"aoc2023/util"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type row struct {
	condition string
	springs   []int
}

func main() {
	input := util.ReadDay(12)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	rows := parseInput(input)

	// fmt.Printf("%+v\n", rows)

	count := 0

	for _, row := range rows {
		c := possibleArrangements("", row)
		// fmt.Println(c)
		count += c
	}

	return count
}

func part2(input []string) int {
	originalRows := parseInput(input)
	rows := []row{}
	for _, row := range originalRows {
		rows = append(rows, expandRow(row))
	}

	// fmt.Printf("%+v\n", rows)

	count := 0

	for _, row := range rows {
		c := possibleArrangements("", row)
		fmt.Println(c)
		count += c
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

func possibleArrangements(prefix string, r row) int {
	i := len(prefix)

	for i < len(r.condition) && r.condition[i] != '?' {
		prefix = prefix + string(r.condition[i])
		i++
	}

	if i == len(r.condition) {
		if isValid(prefix, r.springs) {
			return 1
		} else {
			return 0
		}
	}

	s := calculateSprings(prefix)
	if len(s) > 1 {
		if len(s) > len(r.springs) || !reflect.DeepEqual(s[:len(s)-1], r.springs[:len(s)-1]) {
			return 0
		}
	}

	return possibleArrangements(prefix+"#", r) + possibleArrangements(prefix+".", r)
}

func isValid(condition string, springs []int) bool {
	tmp := calculateSprings(condition)
	// fmt.Printf("%s %+v %+v\n", condition, tmp, springs)
	return reflect.DeepEqual(tmp, springs)
}

func calculateSprings(condition string) []int {
	results := []int{}

	count := 0
	for i, curr := range condition {
		prev := byte('.')
		if i > 0 {
			prev = condition[i-1]
		}
		if curr == '#' {
			count++
		} else if prev == '#' {
			results = append(results, count)
			count = 0
		}
	}
	if count != 0 {
		results = append(results, count)
	}

	return results
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
