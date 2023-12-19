package main

import (
	"aoc2023/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type part struct {
	x int
	m int
	a int
	s int
}

type rule struct {
	rating   string
	op       string
	value    int
	workflow string
}

type ratingRange struct {
	min int
	max int
}

func (p part) getRating(r string) int {
	switch r {
	case "x":
		return p.x
	case "m":
		return p.m
	case "a":
		return p.a
	case "s":
		return p.s
	default:
		return -1
	}
}

func (r rule) apply(p part) (string, bool) {
	if r.op == "" {
	}
	switch r.op {
	case "<":
		if p.getRating(r.rating) < r.value {
			return r.workflow, true
		}
		return "", false
	case ">":
		if p.getRating(r.rating) > r.value {
			return r.workflow, true
		}
		return "", false
	default:
		return r.workflow, true
	}
}

func (rr ratingRange) count() int {
	return rr.max - rr.min + 1
}

func (rr ratingRange) lessThan(val int) (ratingRange, ratingRange) {
	if val <= rr.min {
		return ratingRange{0, -1}, rr
	}
	if val > rr.max {
		return rr, ratingRange{0, -1}
	}
	return ratingRange{rr.min, val - 1}, ratingRange{val, rr.max}
}

func (rr ratingRange) greaterThan(val int) (ratingRange, ratingRange) {
	if val >= rr.max {
		return ratingRange{0, -1}, rr
	}
	if val < rr.min {
		return rr, ratingRange{0, -1}
	}
	return ratingRange{val + 1, rr.max}, ratingRange{rr.min, val}
}

func main() {
	input := util.ReadDay(19)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	result := 0

	workflows, parts := parseInput(input)

	for _, p := range parts {
		if isAccepted("in", p, workflows) {
			result += p.x + p.m + p.a + p.s
		}
	}

	return result
}

func part2(input []string) int {
	workflows, _ := parseInput(input)
	return countCombinations(ratingRange{1, 4000}, ratingRange{1, 4000}, ratingRange{1, 4000}, ratingRange{1, 4000}, "in", workflows)
}

func parseInput(input []string) (workflows map[string][]rule, parts []part) {
	workflows = make(map[string][]rule)

	i := 0
	for ; input[i] != ""; i++ {
		parseWorkflow(input[i], workflows)
	}

	i++
	for ; i < len(input); i++ {
		parts = append(parts, parsePart(input[i]))
	}

	return
}

func parseWorkflow(w string, ws map[string][]rule) {
	var re = regexp.MustCompile(`^(.+)\{(.*)\}$`)
	matches := re.FindStringSubmatch(w)
	for _, r := range strings.Split(matches[2], ",") {
		ws[matches[1]] = append(ws[matches[1]], parseRule(r))
	}
}

func parseRule(r string) rule {
	var re = regexp.MustCompile(`^(([xmas])([<>])(\d+):)?(.*)$`)
	matches := re.FindStringSubmatch(r)
	var value int
	if len(matches[4]) > 0 {
		value, _ = strconv.Atoi(matches[4])
	}
	return rule{matches[2], matches[3], value, matches[5]}
}

func parsePart(p string) part {
	var re = regexp.MustCompile(`^\{x=(\d+),m=(\d+),a=(\d+),s=(\d+)\}$`)
	matches := re.FindStringSubmatch(p)
	x, _ := strconv.Atoi(matches[1])
	m, _ := strconv.Atoi(matches[2])
	a, _ := strconv.Atoi(matches[3])
	s, _ := strconv.Atoi(matches[4])
	return part{x, m, a, s}
}

func isAccepted(w string, p part, workflows map[string][]rule) bool {
	for _, r := range workflows[w] {
		n, ok := r.apply(p)
		if !ok {
			continue
		}
		switch n {
		case "A":
			return true
		case "R":
			return false
		default:
			return isAccepted(n, p, workflows)
		}
	}

	return false
}

func countCombinations(xRange, mRange, aRange, sRange ratingRange, w string, workflows map[string][]rule) int {
	if w == "A" {
		return xRange.count() * mRange.count() * aRange.count() * sRange.count()
	}
	if w == "R" {
		return 0
	}

	result := 0
	var x, m, a, s ratingRange

	for _, r := range workflows[w] {
		switch r.op {
		case "<":
			switch r.rating {
			case "x":
				x, xRange = xRange.lessThan(r.value)
				result += countCombinations(x, mRange, aRange, sRange, r.workflow, workflows)
			case "m":
				m, mRange = mRange.lessThan(r.value)
				result += countCombinations(xRange, m, aRange, sRange, r.workflow, workflows)
			case "a":
				a, aRange = aRange.lessThan(r.value)
				result += countCombinations(xRange, mRange, a, sRange, r.workflow, workflows)
			case "s":
				s, sRange = sRange.lessThan(r.value)
				result += countCombinations(xRange, mRange, aRange, s, r.workflow, workflows)
			}
		case ">":
			switch r.rating {
			case "x":
				x, xRange = xRange.greaterThan(r.value)
				result += countCombinations(x, mRange, aRange, sRange, r.workflow, workflows)
			case "m":
				m, mRange = mRange.greaterThan(r.value)
				result += countCombinations(xRange, m, aRange, sRange, r.workflow, workflows)
			case "a":
				a, aRange = aRange.greaterThan(r.value)
				result += countCombinations(xRange, mRange, a, sRange, r.workflow, workflows)
			case "s":
				s, sRange = sRange.greaterThan(r.value)
				result += countCombinations(xRange, mRange, aRange, s, r.workflow, workflows)
			}
		default:
			result += countCombinations(xRange, mRange, aRange, sRange, r.workflow, workflows)
			break
		}
	}

	return result
}
