package main

import (
	"aoc2023/util"
	"fmt"
)

type loc struct {
	r int
	c int
}

type con struct {
	a loc
	b loc
}

func (c con) other(l loc) loc {
	if c.a == l {
		return c.b
	}
	if c.b == l {
		return c.a
	}
	return loc{-1, -1}
}

func (c con) contains(l loc) bool {
	return c.a == l || c.b == l
}

func (l loc) add(o loc) loc {
	return loc{l.r + o.r, l.c + o.c}
}

func main() {
	input := util.ReadDay(10)

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	m, start := parseInput(input)
	var next loc
	curr := start
	steps := 0

	for _, cardinal := range []loc{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		l := curr.add(cardinal)
		if c, ok := m[l]; ok && c.contains(curr) {
			next = l
			break
		}
	}

	for next != start {
		steps++
		curr, next = next, m[next].other(curr)
	}

	return (steps + 1) / 2
}

func part2(input []string) int {
	m, start := parseInput(input)
	path := map[loc]bool{}
	curr := start

	found := []loc{}
	for _, cardinal := range []loc{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
		l := curr.add(cardinal)
		if c, ok := m[l]; ok && c.contains(curr) {
			found = append(found, l)
		}
	}
	m[curr] = con{found[0], found[1]}
	next := found[0]

	path[curr] = hasSouth(m, curr)

	for next != start {
		curr, next = next, m[next].other(curr)
		path[curr] = hasSouth(m, curr)
	}

	result := 0

	for r := 0; r < len(input); r++ {
		southCount := 0
		for c := 0; c < len(input[r]); c++ {
			l := loc{r, c}
			hasSouth, inPath := path[l]
			if hasSouth {
				southCount++
			}
			if !inPath && southCount%2 == 1 {
				result++
			}
		}
	}

	return result
}

func hasSouth(m map[loc]con, l loc) bool {
	south := l.add(loc{1, 0})
	c1 := m[l]
	if c2, ok := m[south]; ok && c2.contains(l) && c1.contains(south) {
		return true
	}
	return false
}

func parseInput(input []string) (map[loc]con, loc) {
	m := map[loc]con{}
	var start loc

	for r := 0; r < len(input); r++ {
		for c := 0; c < len(input[r]); c++ {
			curr := loc{r, c}
			switch input[r][c] {
			case '|':
				m[curr] = con{loc{r - 1, c}, loc{r + 1, c}}
			case '-':
				m[curr] = con{loc{r, c - 1}, loc{r, c + 1}}
			case 'L':
				m[curr] = con{loc{r - 1, c}, loc{r, c + 1}}
			case 'J':
				m[curr] = con{loc{r - 1, c}, loc{r, c - 1}}
			case '7':
				m[curr] = con{loc{r + 1, c}, loc{r, c - 1}}
			case 'F':
				m[curr] = con{loc{r + 1, c}, loc{r, c + 1}}
			case 'S':
				start = curr
			}
		}
	}

	return m, start
}
