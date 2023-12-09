package main

import (
	"aoc2023/util"
	"fmt"
	"regexp"
	"strings"
)

type node struct {
	l string
	r string
}

func main() {
	input := util.ReadDay(8)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	instr, m := parseInput(input)

	count := 0
	location := "AAA"

	for {
		for _, d := range strings.Split(instr, "") {
			if d == "L" {
				location = m[location].l
			} else {
				location = m[location].r
			}
			count++
			if location == "ZZZ" {
				return count
			}
		}
	}
}

// This logic assumes every path is a loop of equal length, even the first loop.
// This matches my data, but does not seem like a general solution
func part2(input []string) int {
	instr, m := parseInput(input)

	count := 0
	locations := []string{}
	for l, _ := range m {
		if l[2] == byte('A') {
			locations = append(locations, l)
		}
	}

	loops := []int{}

outer:
	for {
		for _, d := range strings.Split(instr, "") {
			nextLocations := []string{}
			for _, location := range locations {
				if d == "L" {
					location = m[location].l
				} else {
					location = m[location].r
				}
				if location[2] != byte('Z') {
					nextLocations = append(nextLocations, location)
				} else {
					loops = append(loops, count+1)
				}
			}
			count++
			locations = nextLocations
			if len(locations) == 0 {
				break outer
			}
		}
	}

	return lcm(loops)
}

func processNode(m map[string]node, line string) {
	re := regexp.MustCompile(`^(...) = \((...), (...)\)$`)
	matches := re.FindStringSubmatch(line)
	m[matches[1]] = node{matches[2], matches[3]}
}

func parseInput(input []string) (string, map[string]node) {
	m := make(map[string]node, 0)

	for i := 2; i < len(input); i++ {
		processNode(m, input[i])
	}

	return input[0], m
}

// Based on https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(integers []int) int {
	result := integers[0] * integers[1] / gcd(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = result * integers[i] / gcd(result, integers[i])
	}

	return result
}
